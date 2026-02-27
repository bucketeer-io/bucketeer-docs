package guardrails

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// Output guardrail limits based on design.md section 4.2
const (
	// MaxOutputSize is the maximum output size per file (64KB)
	MaxOutputSize = 64 * 1024

	// OpenTag is the opening tag for updated document content
	OpenTag = "<updated_document>"

	// CloseTag is the closing tag for updated document content
	CloseTag = "</updated_document>"
)

// Output validation errors
var (
	ErrMissingDocumentTags = errors.New("missing <updated_document> tags in output")
	ErrOutputTooLarge      = errors.New("output content exceeds maximum size")
	ErrInvalidMarkdown     = errors.New("invalid markdown syntax detected")
	ErrEmptyContent        = errors.New("extracted document content is empty")
)

// OutputGuardrails provides output validation for AI-generated content
type OutputGuardrails struct {
	MaxOutputSize int
}

// NewOutputGuardrails creates a new OutputGuardrails with default settings
func NewOutputGuardrails() *OutputGuardrails {
	return &OutputGuardrails{
		MaxOutputSize: MaxOutputSize,
	}
}

// Validate performs all output validation checks on AI-generated content
func (g *OutputGuardrails) Validate(output string) error {
	// 1. Validate <updated_document> tags are present
	if err := g.validateDocumentTags(output); err != nil {
		return err
	}

	// 2. Extract document content
	content := ExtractDocumentContent(output)
	if content == "" {
		return ErrEmptyContent
	}

	// 3. Check size limit
	if len(content) > g.MaxOutputSize {
		return fmt.Errorf("%w: %d bytes (max %d)", ErrOutputTooLarge, len(content), g.MaxOutputSize)
	}

	// 4. Basic markdown syntax validation
	if err := g.validateMarkdown(content); err != nil {
		return err
	}

	return nil
}

// validateDocumentTags checks for the presence of required document tags
func (g *OutputGuardrails) validateDocumentTags(output string) error {
	if !strings.Contains(output, OpenTag) || !strings.Contains(output, CloseTag) {
		return ErrMissingDocumentTags
	}
	return nil
}

// validateMarkdown performs basic markdown syntax validation
func (g *OutputGuardrails) validateMarkdown(content string) error {
	// Check for unbalanced code blocks (triple backticks)
	codeBlockCount := strings.Count(content, "```")
	if codeBlockCount%2 != 0 {
		return fmt.Errorf("%w: unbalanced code blocks (found %d triple backticks)", ErrInvalidMarkdown, codeBlockCount)
	}

	// Check for frontmatter if present (must be properly closed)
	trimmed := strings.TrimSpace(content)
	if strings.HasPrefix(trimmed, "---") {
		rest := trimmed[3:]
		if !strings.Contains(rest, "---") {
			return fmt.Errorf("%w: unclosed frontmatter", ErrInvalidMarkdown)
		}
	}

	// Check for severely malformed headers (# without space before text)
	// This is a common AI generation issue
	malformedHeaderRegex := regexp.MustCompile(`(?m)^#{1,6}[^\s#]`)
	if malformedHeaderRegex.MatchString(content) {
		return fmt.Errorf("%w: malformed header (missing space after #)", ErrInvalidMarkdown)
	}

	return nil
}

// ExtractDocumentContent extracts content from within <updated_document> tags
func ExtractDocumentContent(output string) string {
	startIdx := strings.Index(output, OpenTag)
	if startIdx == -1 {
		return ""
	}
	startIdx += len(OpenTag)

	endIdx := strings.Index(output, CloseTag)
	if endIdx == -1 || startIdx >= endIdx {
		return ""
	}

	return strings.TrimSpace(output[startIdx:endIdx])
}

// TransformNonASCII normalizes Unicode and replaces typographic punctuation.
// Returns transformed content and list of warnings when changes occur.
func TransformNonASCII(content string) (string, []string) {
	var warnings []string
	normalized := norm.NFKC.String(content)
	if normalized != content {
		warnings = append(warnings, "Normalized Unicode with NFKC")
	}
	return normalized, warnings
}

// PostProcess transforms AI output to fix common issues.
// Returns processed content and list of transformations applied.
func (g *OutputGuardrails) PostProcess(content string) (string, []string) {
	var allWarnings []string

	// 1. Transform non-ASCII punctuation
	result, warnings := TransformNonASCII(content)
	allWarnings = append(allWarnings, warnings...)

	// 2. Check for MDX compatibility issues
	mdxWarnings := CheckMDXCompatibility(result)
	allWarnings = append(allWarnings, mdxWarnings...)

	// 3. Ensure trailing newline (standard for text files)
	result = normalizeTrailingNewline(result)

	return result, allWarnings
}

// normalizeTrailingNewline ensures content ends with exactly one newline.
func normalizeTrailingNewline(content string) string {
	// Trim trailing whitespace and newlines, then add exactly one newline
	trimmed := strings.TrimRight(content, " \t\n\r")
	return trimmed + "\n"
}

// CheckMDXCompatibility checks for patterns that may cause MDX parsing issues.
// Returns warnings for any detected issues (does not modify content).
func CheckMDXCompatibility(content string) []string {
	var warnings []string

	// Skip content inside code blocks for analysis
	// Simple heuristic: check lines not inside triple-backtick blocks
	lines := strings.Split(content, "\n")
	inCodeBlock := false

	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCodeBlock = !inCodeBlock
			continue
		}

		if inCodeBlock {
			continue
		}

		// Check for problematic patterns outside code blocks
		if strings.Contains(line, "<=") || strings.Contains(line, ">=") {
			// Exclude if inside inline code (backticks)
			if !isInsideInlineCode(line, "<=") && strings.Contains(line, "<=") {
				warnings = append(warnings, fmt.Sprintf("Line %d: '<=' may break MDX - use '&lt;=' instead", i+1))
			}
			if !isInsideInlineCode(line, ">=") && strings.Contains(line, ">=") {
				warnings = append(warnings, fmt.Sprintf("Line %d: '>=' may break MDX - use '&gt;=' instead", i+1))
			}
		}
	}

	return warnings
}

// isInsideInlineCode checks if a pattern appears inside backtick-enclosed code.
func isInsideInlineCode(line, pattern string) bool {
	idx := strings.Index(line, pattern)
	if idx == -1 {
		return false
	}

	// Count backticks before the pattern
	before := line[:idx]
	backtickCount := strings.Count(before, "`")

	// Odd count means we're inside inline code
	return backtickCount%2 == 1
}
