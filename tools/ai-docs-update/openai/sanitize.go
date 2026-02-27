// Package openai provides OpenAI API client and prompt handling for AI-driven documentation updates.
package openai

import (
	"strings"
)

// SanitizedContext holds sanitized Issue/PR context for safe prompt inclusion.
type SanitizedContext struct {
	IssueTitle string
	IssueBody  string
	PRTitle    string
	PRBody     string
	Diff       string
}

// Constants for sanitization limits
const (
	MaxTitleLen = 200
	MaxBodyLen  = 10 * 1024 // 10KB
	MaxDiffLen  = 50000     // 50KB
)

// NewSanitizedContext creates a SanitizedContext with all inputs properly sanitized.
// This protects against prompt injection attacks by escaping XML tags and code blocks.
func NewSanitizedContext(issueTitle, issueBody, prTitle, prBody, diff string) SanitizedContext {
	return SanitizedContext{
		IssueTitle: SanitizeForPrompt(issueTitle, MaxTitleLen),
		IssueBody:  SanitizeForPrompt(issueBody, MaxBodyLen),
		PRTitle:    SanitizeForPrompt(prTitle, MaxTitleLen),
		PRBody:     SanitizeForPrompt(prBody, MaxBodyLen),
		Diff:       SanitizeForPrompt(diff, MaxDiffLen),
	}
}

// SanitizeForPrompt sanitizes input for safe prompt inclusion.
// Uses allow-list approach: escapes dangerous patterns rather than trying to remove them.
//
// Protection measures:
// 1. Length truncation to prevent context overflow
// 2. XML tag escaping to prevent prompt structure manipulation
// 3. Code block escaping to prevent Markdown injection
func SanitizeForPrompt(input string, maxLen int) string {
	// 1. Length limit
	if len(input) > maxLen {
		input = input[:maxLen] + "...[truncated]"
	}

	// 2. XML tag escaping (prevent prompt structure breaking)
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")

	// 3. Code block escaping (prevent Markdown injection)
	input = strings.ReplaceAll(input, "```", "` ` `")

	return input
}
