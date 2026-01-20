// Package guardrails provides input and output validation for the AI docs update tool.
package guardrails

import (
	"errors"
	"fmt"
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
	ErrDiffTooLarge      = errors.New("diff size exceeds maximum allowed")
	ErrTooManyFiles      = errors.New("too many changed files")
	ErrFileTooLarge      = errors.New("file exceeds maximum lines")
	ErrDocTooLarge       = errors.New("document content exceeds maximum size")
	ErrTokenLimitExceeded = errors.New("total token count exceeds limit")
	ErrIssueBodyTooLarge = errors.New("issue body exceeds maximum size")
	ErrPRBodyTooLarge    = errors.New("PR body exceeds maximum size")
)

// DiffFile represents a single file in a diff
type DiffFile struct {
	Path      string
	LineCount int
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

		if len(pr.Diff) > g.MaxDiffSizeBytes {
			return fmt.Errorf("%w: %d bytes (max %d)", ErrDiffTooLarge, len(pr.Diff), g.MaxDiffSizeBytes)
		}
	}

	return nil
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
	// Simple prefix matching for now
	// TODO: Implement full glob pattern matching if needed
	if len(pattern) == 0 {
		return false
	}

	// Handle wildcard suffix (e.g., "proto/**")
	if len(pattern) >= 2 && pattern[len(pattern)-2:] == "**" {
		prefix := pattern[:len(pattern)-2]
		return len(path) >= len(prefix) && path[:len(prefix)] == prefix
	}

	return path == pattern
}
