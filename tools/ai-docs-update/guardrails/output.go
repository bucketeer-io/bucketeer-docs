package guardrails

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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

// Non-ASCII punctuation replacements
var nonASCIIReplacements = map[rune]string{
	// Dashes
	'\u2010': "-",  // hyphen → ASCII hyphen
	'\u2011': "-",  // non-breaking hyphen → ASCII hyphen
	'\u2012': "-",  // figure dash → ASCII hyphen
	'\u2013': "-",  // en-dash → ASCII hyphen
	'\u2014': "--", // em-dash → double hyphen
	'\u2015': "--", // horizontal bar → double hyphen

	// Quotes
	'\u2018': "'",   // left single quote → apostrophe
	'\u2019': "'",   // right single quote → apostrophe
	'\u201C': "\"",  // left double quote → quote
	'\u201D': "\"",  // right double quote → quote
	'\u2026': "...", // ellipsis → three dots

	// Math symbols
	'\u2264': "<=", // ≤ less-than or equal → <=
	'\u2265': ">=", // ≥ greater-than or equal → >=
	'\u2260': "!=", // ≠ not equal → !=
	'\u00D7': "x",  // × multiplication sign → x
	'\u00F7': "/",  // ÷ division sign → /

	// Special spaces
	'\u00A0': " ", // NO-BREAK SPACE → regular space
	'\u202F': " ", // NARROW NO-BREAK SPACE → regular space
	'\u2009': " ", // THIN SPACE → regular space
	'\u200B': "",  // ZERO WIDTH SPACE → remove
}

// TransformNonASCII replaces non-ASCII punctuation with ASCII equivalents.
// Returns transformed content and list of unique replacements made.
func TransformNonASCII(content string) (string, []string) {
	var result strings.Builder
	seen := make(map[rune]bool)

	for _, r := range content {
		if replacement, found := nonASCIIReplacements[r]; found {
			result.WriteString(replacement)
			seen[r] = true
		} else {
			result.WriteRune(r)
		}
	}

	// Build warnings for unique replacements only
	var warnings []string
	for r := range seen {
		warnings = append(warnings, fmt.Sprintf("Replaced non-ASCII '%c' (U+%04X) with '%s'", r, r, nonASCIIReplacements[r]))
	}

	return result.String(), warnings
}

// PostProcess transforms AI output to fix common issues.
// Returns processed content and list of transformations applied.
func (g *OutputGuardrails) PostProcess(content string) (string, []string) {
	var allWarnings []string

	// 1. Transform non-ASCII punctuation
	result, warnings := TransformNonASCII(content)
	allWarnings = append(allWarnings, warnings...)

	// 2. Ensure trailing newline (standard for text files)
	result = normalizeTrailingNewline(result)

	return result, allWarnings
}

// normalizeTrailingNewline ensures content ends with exactly one newline.
func normalizeTrailingNewline(content string) string {
	// Trim trailing whitespace and newlines, then add exactly one newline
	trimmed := strings.TrimRight(content, " \t\n\r")
	return trimmed + "\n"
}
