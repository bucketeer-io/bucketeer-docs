// Package guardrails provides input and output validation for the AI docs update tool.
package guardrails

import (
	"errors"
	"fmt"
	"strings"
)

// Input guardrail limits based on design.md section 4.1
const (
	// MaxDiffSizeBytes is the maximum allowed size for code diffs (50000 bytes)
	// Aligned with openai/sanitize.go MaxDiffLen for consistency
	MaxDiffSizeBytes = 50000

	// MaxChangedFiles is the maximum number of changed files allowed
	MaxChangedFiles = 30

	// MaxLinesPerFile is the maximum number of lines allowed per file
	MaxLinesPerFile = 1000

	// MaxDocContentBytes is the maximum size for document content (32KB)
	MaxDocContentBytes = 32 * 1024

	// MaxInputTokens is the maximum token limit for GPT-4o (with margin from 128K limit)
	MaxInputTokens = 100000

	// MaxIssueBodyLen is the maximum length for issue body (10KB)
	MaxIssueBodyLen = 10 * 1024

	// MaxPRBodyLen is the maximum length for PR body (10KB)
	MaxPRBodyLen = 10 * 1024

	// MaxTitleLen is the maximum length for issue/PR titles
	MaxTitleLen = 200
)

// Input validation errors
var (
	ErrDiffTooLarge       = errors.New("diff size exceeds maximum allowed")
	ErrTooManyFiles       = errors.New("too many changed files")
	ErrFileTooLarge       = errors.New("file exceeds maximum lines")
	ErrDocTooLarge        = errors.New("document content exceeds maximum size")
	ErrTokenLimitExceeded = errors.New("total token count exceeds limit")
	ErrIssueBodyTooLarge  = errors.New("issue body exceeds maximum size")
	ErrPRBodyTooLarge     = errors.New("PR body exceeds maximum size")
)

// DiffFile represents a single file in a diff
type DiffFile struct {
	Path      string
	LineCount int
	Additions int    // Number of added lines
	Deletions int    // Number of deleted lines
	Content   string // Raw diff content for this file
}

// Diff represents a code diff with metadata
type Diff struct {
	TotalSize int
	Files     []DiffFile
}

// IssueContext represents the issue context for validation
type IssueContext struct {
	Title string
	Body  string
}

// PRContext represents the PR context for validation
type PRContext struct {
	Title string
	Body  string
	Diff  string
}

// InputGuardrails provides input validation for the AI docs update tool
type InputGuardrails struct {
	MaxDiffSizeBytes   int
	MaxChangedFiles    int
	MaxLinesPerFile    int
	MaxDocContentBytes int
	MaxInputTokens     int
	MaxIssueBodyLen    int
	MaxPRBodyLen       int
	AllowedSourcePaths []string
}

// NewInputGuardrails creates a new InputGuardrails with default limits
func NewInputGuardrails() *InputGuardrails {
	return &InputGuardrails{
		MaxDiffSizeBytes:   MaxDiffSizeBytes,
		MaxChangedFiles:    MaxChangedFiles,
		MaxLinesPerFile:    MaxLinesPerFile,
		MaxDocContentBytes: MaxDocContentBytes,
		MaxInputTokens:     MaxInputTokens,
		MaxIssueBodyLen:    MaxIssueBodyLen,
		MaxPRBodyLen:       MaxPRBodyLen,
		AllowedSourcePaths: []string{}, // Empty means all paths allowed
	}
}

// ValidateDiff validates the diff against size and file count limits
func (g *InputGuardrails) ValidateDiff(diff *Diff) error {
	if diff == nil {
		return nil
	}

	if diff.TotalSize > g.MaxDiffSizeBytes {
		return fmt.Errorf("%w: %d bytes (max %d)", ErrDiffTooLarge, diff.TotalSize, g.MaxDiffSizeBytes)
	}

	if len(diff.Files) > g.MaxChangedFiles {
		return fmt.Errorf("%w: %d files (max %d)", ErrTooManyFiles, len(diff.Files), g.MaxChangedFiles)
	}

	for _, file := range diff.Files {
		if file.LineCount > g.MaxLinesPerFile {
			return fmt.Errorf("%w: %s has %d lines (max %d)", ErrFileTooLarge, file.Path, file.LineCount, g.MaxLinesPerFile)
		}
	}

	return nil
}

// ValidateDocContent validates that a document's content is within size limits
func (g *InputGuardrails) ValidateDocContent(content string) error {
	if len(content) > g.MaxDocContentBytes {
		return fmt.Errorf("%w: %d bytes (max %d)", ErrDocTooLarge, len(content), g.MaxDocContentBytes)
	}
	return nil
}

// ValidateTokenLimit checks if the combined context exceeds the token limit
func (g *InputGuardrails) ValidateTokenLimit(prContext, docContent string) error {
	total := estimateTokens(prContext) + estimateTokens(docContent)
	if total > g.MaxInputTokens {
		return fmt.Errorf("%w: %d tokens (max %d)", ErrTokenLimitExceeded, total, g.MaxInputTokens)
	}
	return nil
}

// ValidateContext validates the issue and PR context
func (g *InputGuardrails) ValidateContext(issue *IssueContext, pr *PRContext) error {
	if issue != nil {
		if len(issue.Body) > g.MaxIssueBodyLen {
			return fmt.Errorf("%w: %d bytes (max %d)", ErrIssueBodyTooLarge, len(issue.Body), g.MaxIssueBodyLen)
		}
	}

	if pr != nil {
		if len(pr.Body) > g.MaxPRBodyLen {
			return fmt.Errorf("%w: %d bytes (max %d)", ErrPRBodyTooLarge, len(pr.Body), g.MaxPRBodyLen)
		}
		// Note: Diff size is NOT validated here - large diffs are summarized instead
		// See SummarizeLargeDiff for handling
	}

	return nil
}

// SummarizeLargeDiff creates a summary for diffs exceeding the size limit.
// Returns the original diff if it's within limits, or a summary otherwise.
func SummarizeLargeDiff(diff string, maxSize int) string {
	if len(diff) <= maxSize {
		return diff
	}

	// Parse the diff to extract file information
	parsed := ParseDiff(diff)

	var summary strings.Builder
	summary.WriteString("## Diff Summary (original too large)\n\n")
	summary.WriteString(fmt.Sprintf("Total: %d files changed, %d bytes\n\n", len(parsed.Files), parsed.TotalSize))
	summary.WriteString("### Changed Files:\n")

	for _, f := range parsed.Files {
		summary.WriteString(fmt.Sprintf("- %s (+%d/-%d lines)\n", f.Path, f.Additions, f.Deletions))
	}

	// Include truncated content of key files (prioritize non-test, non-vendor files)
	summary.WriteString("\n### Key Changes (truncated):\n")
	remainingBytes := maxSize - summary.Len() - 100 // Reserve space

	for _, f := range parsed.Files {
		// Skip test files and vendor
		if isTestOrVendorFile(f.Path) {
			continue
		}

		if remainingBytes <= 0 {
			break
		}

		// Include first portion of each file's diff
		content := f.Content
		if len(content) > remainingBytes/len(parsed.Files) {
			content = content[:remainingBytes/len(parsed.Files)] + "\n...[truncated]"
		}

		summary.WriteString(fmt.Sprintf("\n#### %s\n```diff\n%s\n```\n", f.Path, content))
		remainingBytes -= len(content) + 50
	}

	return summary.String()
}

// isTestOrVendorFile checks if a file path is a test or vendor file
func isTestOrVendorFile(path string) bool {
	lowerPath := strings.ToLower(path)
	return strings.Contains(lowerPath, "_test.") ||
		strings.Contains(lowerPath, "/test/") ||
		strings.Contains(lowerPath, "/vendor/") ||
		strings.Contains(lowerPath, "/node_modules/") ||
		strings.HasSuffix(lowerPath, ".test.ts") ||
		strings.HasSuffix(lowerPath, ".test.tsx") ||
		strings.HasSuffix(lowerPath, ".spec.ts") ||
		strings.HasSuffix(lowerPath, ".spec.tsx")
}

// IsAllowedPath checks if a file path is in the allowed source paths
// If AllowedSourcePaths is empty, all paths are allowed
func (g *InputGuardrails) IsAllowedPath(path string) bool {
	if len(g.AllowedSourcePaths) == 0 {
		return true
	}

	for _, allowed := range g.AllowedSourcePaths {
		if matchPath(path, allowed) {
			return true
		}
	}
	return false
}

// estimateTokens provides a simple token estimation
// Rule of thumb: 1 token ≈ 4 characters for English text
func estimateTokens(text string) int {
	return len(text) / 4
}

// matchPath performs simple path matching with glob-like patterns
func matchPath(path, pattern string) bool {
	if pattern == "" {
		return false
	}

	// Handle wildcard suffix (e.g., "proto/**")
	if strings.HasSuffix(pattern, "**") {
		prefix := pattern[:len(pattern)-2]
		return strings.HasPrefix(path, prefix)
	}

	return path == pattern
}

// ParseDiff parses a unified diff and extracts file information.
// Uses simple regex to count files and estimate line changes.
func ParseDiff(diffContent string) *Diff {
	if diffContent == "" {
		return nil
	}

	diff := &Diff{
		TotalSize: len(diffContent),
		Files:     []DiffFile{},
	}

	lines := strings.Split(diffContent, "\n")
	var currentFile *DiffFile
	var contentBuilder strings.Builder
	additions, deletions := 0, 0

	for _, line := range lines {
		// Detect new file in diff (diff --git a/path b/path)
		if strings.HasPrefix(line, "diff --git ") {
			// Save previous file if exists
			if currentFile != nil {
				currentFile.LineCount = additions + deletions
				currentFile.Additions = additions
				currentFile.Deletions = deletions
				currentFile.Content = contentBuilder.String()
				diff.Files = append(diff.Files, *currentFile)
			}

			// Extract file path from "diff --git a/path b/path"
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				// Remove "a/" or "b/" prefix
				path := strings.TrimPrefix(parts[2], "a/")
				currentFile = &DiffFile{Path: path}
				contentBuilder.Reset()
				additions, deletions = 0, 0
			}
			continue
		}

		// Collect content and count lines
		if currentFile != nil {
			contentBuilder.WriteString(line)
			contentBuilder.WriteString("\n")

			if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
				additions++
			} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
				deletions++
			}
		}
	}

	// Save last file
	if currentFile != nil {
		currentFile.LineCount = additions + deletions
		currentFile.Additions = additions
		currentFile.Deletions = deletions
		currentFile.Content = contentBuilder.String()
		diff.Files = append(diff.Files, *currentFile)
	}

	return diff
}

// SummarizeDiff creates a concise summary of diff for Phase 1.
// Returns a list of changed files with change type indicators.
func SummarizeDiff(diffContent string) string {
	if diffContent == "" {
		return ""
	}

	parsed := ParseDiff(diffContent)
	if parsed == nil || len(parsed.Files) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("Changed files:\n")
	for _, f := range parsed.Files {
		changeType := "modified"
		if f.LineCount > 100 {
			changeType = "heavily modified"
		}
		sb.WriteString(fmt.Sprintf("- %s (%s, ~%d lines)\n", f.Path, changeType, f.LineCount))
	}

	return sb.String()
}
