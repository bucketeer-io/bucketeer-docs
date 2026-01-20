// Package openai provides OpenAI API client and prompt handling for AI-driven documentation updates.
package openai

import (
	"strings"
)

// Sanitization limits for prompt injection protection.
const (
	MaxTitleLen     = 200
	MaxBodyLen      = 10000
	MaxIssueBodyLen = 10000
	MaxDiffLen      = 50000
)

// SanitizedContext holds sanitized Issue/PR context for safe prompt inclusion.
type SanitizedContext struct {
	IssueTitle string
	IssueBody  string
	PRTitle    string
	PRBody     string
	Diff       string
}

// NewSanitizedContext creates a SanitizedContext with all inputs properly sanitized.
// This protects against prompt injection attacks by escaping XML tags and code blocks.
func NewSanitizedContext(issueTitle, issueBody, prTitle, prBody, diff string) SanitizedContext {
	return SanitizedContext{
		IssueTitle: SanitizeForPrompt(issueTitle, MaxTitleLen),
		IssueBody:  SanitizeForPrompt(issueBody, MaxIssueBodyLen),
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

// SanitizeTitle is a convenience function for sanitizing titles.
func SanitizeTitle(title string) string {
	return SanitizeForPrompt(title, MaxTitleLen)
}

// SanitizeBody is a convenience function for sanitizing body content.
func SanitizeBody(body string) string {
	return SanitizeForPrompt(body, MaxBodyLen)
}

// SanitizeIssueBody is a convenience function for sanitizing issue body content.
func SanitizeIssueBody(body string) string {
	return SanitizeForPrompt(body, MaxIssueBodyLen)
}

// SanitizeDiff is a convenience function for sanitizing diff content.
func SanitizeDiff(diff string) string {
	return SanitizeForPrompt(diff, MaxDiffLen)
}
