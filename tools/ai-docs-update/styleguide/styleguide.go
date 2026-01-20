// Package styleguide loads and parses documentation style guidelines.
package styleguide

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// essentialFiles lists the style guide files to extract rules from.
// These contain the most important rules for AI-generated content.
var essentialFiles = []struct {
	filename string
	category string
}{
	{"02-voice-and-tone.md", "Voice and Tone"},
	{"04-language-and-grammar.md", "Language and Grammar"},
	{"06-ui-elements-and-interaction.md", "UI Elements"},
	{"08-code-elements.md", "Code Elements"},
}

// StyleGuide contains extracted style rules organized by category.
type StyleGuide struct {
	Categories []Category
}

// Category groups related style rules.
type Category struct {
	Name  string
	Rules []string
}

// Load reads style guide files from the given directory and extracts essential rules.
// The directory should be the path to docs/contribution-guide/documentation-style/.
func Load(styleDir string) (*StyleGuide, error) {
	sg := &StyleGuide{}

	for _, ef := range essentialFiles {
		filePath := filepath.Join(styleDir, ef.filename)
		rules, err := extractRules(filePath)
		if err != nil {
			// Log warning but continue with other files
			fmt.Fprintf(os.Stderr, "Warning: failed to load %s: %v\n", ef.filename, err)
			continue
		}

		if len(rules) > 0 {
			sg.Categories = append(sg.Categories, Category{
				Name:  ef.category,
				Rules: rules,
			})
		}
	}

	return sg, nil
}

// extractRules reads a markdown file and extracts bullet point rules.
// It filters out "Not recommended" examples and keeps only actionable guidelines.
func extractRules(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rules []string
	scanner := bufio.NewScanner(file)
	inNotRecommended := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Track whether we're in a "Not recommended" section to skip bad examples
		inNotRecommended = updateSectionState(trimmed, inNotRecommended)
		if inNotRecommended {
			continue
		}

		// Extract bullet point rules (lines starting with - or *)
		rule, isBullet := extractBulletContent(trimmed)
		if !isBullet {
			continue
		}

		if shouldSkipRule(rule) {
			continue
		}

		// Clean up italic markers
		rule = strings.ReplaceAll(rule, "*", "")

		// Limit rule length for prompt efficiency
		if len(rule) > 150 {
			rule = rule[:147] + "..."
		}

		rules = append(rules, rule)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Deduplicate and limit rules per category
	return deduplicateRules(rules, 8), nil
}

// updateSectionState determines if we're entering or leaving a "Not recommended" section.
func updateSectionState(line string, currentState bool) bool {
	// Enter "Not recommended" section
	if strings.Contains(line, "Not recommended") || strings.Contains(line, ":x:") {
		return true
	}
	// Exit when hitting "Recommended", thumbs up emoji, or new heading
	if strings.Contains(line, "Recommended") ||
		strings.Contains(line, ":+1:") ||
		strings.HasPrefix(line, "##") {
		return false
	}
	return currentState
}

// extractBulletContent returns the content of a bullet point line and whether it is a bullet.
func extractBulletContent(line string) (string, bool) {
	if content, found := strings.CutPrefix(line, "- "); found {
		return content, true
	}
	if content, found := strings.CutPrefix(line, "* "); found {
		return content, true
	}
	return "", false
}

// shouldSkipRule returns true if the rule should be skipped (example lines, empty, or too short).
func shouldSkipRule(rule string) bool {
	if rule == "" || len(rule) < 10 {
		return true
	}
	if strings.HasPrefix(rule, "The ") || strings.HasPrefix(rule, "This ") {
		return true
	}
	return strings.Contains(rule, "example")
}

// deduplicateRules removes duplicate rules and limits the count.
func deduplicateRules(rules []string, maxCount int) []string {
	seen := make(map[string]bool)
	var unique []string

	for _, rule := range rules {
		normalized := strings.ToLower(rule)
		if !seen[normalized] {
			seen[normalized] = true
			unique = append(unique, rule)
			if len(unique) >= maxCount {
				break
			}
		}
	}

	return unique
}

// Format returns the style guide as prompt-ready markdown text.
func (sg *StyleGuide) Format() string {
	if sg == nil || len(sg.Categories) == 0 {
		return defaultStyleGuide()
	}

	var sb strings.Builder
	for _, cat := range sg.Categories {
		if len(cat.Rules) == 0 {
			continue
		}

		sb.WriteString("### ")
		sb.WriteString(cat.Name)
		sb.WriteString("\n")

		for _, rule := range cat.Rules {
			sb.WriteString("- ")
			sb.WriteString(rule)
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	result := sb.String()
	if result == "" {
		return defaultStyleGuide()
	}

	return result
}

// defaultStyleGuide returns a minimal fallback style guide.
func defaultStyleGuide() string {
	return `### Voice and Tone
- Use conversational, friendly, and professional tone
- Avoid figures of speech and internet slang
- Keep sentences under 25-30 words

### Language and Grammar
- Use active voice: "The team deploys..." not "is deployed by the team"
- Use contractions (don't, it's, you'll) for informal tone
- State goal/condition before instruction

### UI Elements
- Use **bold** for UI elements: buttons, menus, dialogs
- Use sentence case for all UI labels

### Code Elements
- Use backticks for: filenames, class names, method names, HTTP status codes
- Use code blocks for code samples
`
}

// RuleCount returns the total number of rules loaded.
func (sg *StyleGuide) RuleCount() int {
	if sg == nil {
		return 0
	}
	count := 0
	for _, cat := range sg.Categories {
		count += len(cat.Rules)
	}
	return count
}
