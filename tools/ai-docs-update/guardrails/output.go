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
	ErrForbiddenPattern    = errors.New("forbidden pattern found in output")
	ErrInvalidMarkdown     = errors.New("invalid markdown syntax detected")
	ErrEmptyContent        = errors.New("extracted document content is empty")
)

// Default forbidden patterns that should not appear in documentation
var defaultForbiddenPatterns = []string{
	"OPENAI_API_KEY",
	"SECRET_KEY",
	"API_SECRET",
	"password=",
	"BEGIN RSA PRIVATE KEY",
	"BEGIN PRIVATE KEY",
	"BEGIN EC PRIVATE KEY",
	"BEGIN DSA PRIVATE KEY",
	"BEGIN OPENSSH PRIVATE KEY",
}

// OutputGuardrails provides output validation for AI-generated content
type OutputGuardrails struct {
	MaxOutputSize     int
	ForbiddenPatterns []string
}

// NewOutputGuardrails creates a new OutputGuardrails with default settings
func NewOutputGuardrails() *OutputGuardrails {
	return &OutputGuardrails{
		MaxOutputSize:     MaxOutputSize,
		ForbiddenPatterns: defaultForbiddenPatterns,
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

	// 4. Check forbidden patterns
	if err := g.checkForbiddenPatterns(content); err != nil {
		return err
	}

	// 5. Basic markdown syntax validation
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

// checkForbiddenPatterns scans content for any forbidden patterns
func (g *OutputGuardrails) checkForbiddenPatterns(content string) error {
	for _, pattern := range g.ForbiddenPatterns {
		if strings.Contains(content, pattern) {
			return fmt.Errorf("%w: %s", ErrForbiddenPattern, pattern)
		}
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

// AddForbiddenPattern adds a new forbidden pattern to the guardrails
func (g *OutputGuardrails) AddForbiddenPattern(pattern string) {
	g.ForbiddenPatterns = append(g.ForbiddenPatterns, pattern)
}

// SetForbiddenPatterns replaces all forbidden patterns
func (g *OutputGuardrails) SetForbiddenPatterns(patterns []string) {
	g.ForbiddenPatterns = patterns
}

// Placeholder patterns that indicate AI guessing
var placeholderPatterns = map[string]string{
	"X.Y.Z":     "<!-- TODO: Needs confirmation - replace X.Y.Z with actual version -->",
	"N.N.N":     "<!-- TODO: Needs confirmation - replace N.N.N with actual version -->",
	"vX.X.X":    "<!-- TODO: Needs confirmation - replace vX.X.X with actual version -->",
	"[version]": "<!-- TODO: Needs confirmation - specify version number -->",
	"[VERSION]": "<!-- TODO: Needs confirmation - specify version number -->",
	"<version>": "<!-- TODO: Needs confirmation - specify version number -->",
	"TBD":       "<!-- TODO: Needs confirmation - TBD -->",
}

// Non-ASCII punctuation replacements
var nonASCIIReplacements = map[rune]string{
	'\u2010': "-",   // hyphen → ASCII hyphen
	'\u2011': "-",   // non-breaking hyphen → ASCII hyphen
	'\u2012': "-",   // figure dash → ASCII hyphen
	'\u2013': "-",   // en-dash → ASCII hyphen
	'\u2014': "--",  // em-dash → double hyphen
	'\u2015': "--",  // horizontal bar → double hyphen
	'\u2018': "'",   // left single quote → apostrophe
	'\u2019': "'",   // right single quote → apostrophe
	'\u201C': "\"",  // left double quote → quote
	'\u201D': "\"",  // right double quote → quote
	'\u2026': "...", // ellipsis → three dots
	'\u00A0': " ",   // non-breaking space → regular space
}

// TransformPlaceholders converts version placeholders to TODO comments.
// Returns transformed content and list of warnings.
func TransformPlaceholders(content string) (string, []string) {
	var warnings []string
	result := content
	for pattern, todoComment := range placeholderPatterns {
		if strings.Contains(result, pattern) {
			// Add TODO comment after the placeholder
			result = strings.ReplaceAll(result, pattern, pattern+" "+todoComment)
			warnings = append(warnings, fmt.Sprintf("Placeholder '%s' found - added TODO marker", pattern))
		}
	}
	return result, warnings
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

	// 1. Transform placeholders to TODO comments
	result, warnings := TransformPlaceholders(content)
	allWarnings = append(allWarnings, warnings...)

	// 2. Transform non-ASCII punctuation
	result, warnings = TransformNonASCII(result)
	allWarnings = append(allWarnings, warnings...)

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
