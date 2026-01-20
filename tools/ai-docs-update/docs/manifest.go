// Package docs provides functionality for working with documentation files.
package docs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
)

// Manifest holds a list of all documentation files.
type Manifest struct {
	Files []DocFile `json:"files"`
}

// DefaultExcludeDirs lists directories excluded from manifest generation by default.
// These directories contain auto-generated or static content that should not be modified.
var DefaultExcludeDirs = []string{
	"changelog",          // Auto-generated from GitHub release notes
	"contribution-guide", // Static contribution guidelines (not feature-related)
}

// DocFile represents a single documentation file.
type DocFile struct {
	Path        string   `json:"path"`        // Relative path from docs root (e.g., "feature-flags/segments.mdx")
	Title       string   `json:"title"`       // Title from frontmatter
	Description string   `json:"description"` // First paragraph of content
	Tags        []string `json:"tags"`        // Tags from frontmatter for categorization
	Category    string   `json:"category"`    // Inferred category (sdk, feature-flags, etc.)
	Audience    string   `json:"audience"`    // Inferred audience (external-developers, operators, admins)
}

// FrontMatter represents the frontmatter of a documentation file.
type FrontMatter struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Slug        string   `yaml:"slug"`
	Sidebar     string   `yaml:"sidebar_label"`
	Tags        []string `yaml:"tags"`
}

// GenerateManifest scans the docs directory and generates a manifest of all .mdx files.
// Directories in excludeDirs are skipped. If excludeDirs is nil, DefaultExcludeDirs is used.
func GenerateManifest(docsDir string, excludeDirs []string) (*Manifest, error) {
	if excludeDirs == nil {
		excludeDirs = DefaultExcludeDirs
	}

	// Convert to map for O(1) lookup
	excludeSet := make(map[string]bool)
	for _, dir := range excludeDirs {
		excludeSet[dir] = true
	}

	var files []DocFile

	err := filepath.Walk(docsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path for exclusion check
		relPath, _ := filepath.Rel(docsDir, path)

		// Skip excluded directories entirely
		if info.IsDir() {
			// Check if this directory is in exclude list
			topDir := strings.Split(relPath, string(filepath.Separator))[0]
			if excludeSet[topDir] {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process .mdx files
		if !strings.HasSuffix(path, ".mdx") {
			return nil
		}

		docFile, err := parseDocFile(docsDir, path)
		if err != nil {
			// Log warning but continue processing other files
			fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", path, err)
			return nil
		}

		files = append(files, *docFile)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk docs directory: %w", err)
	}

	return &Manifest{Files: files}, nil
}

// parseDocFile parses a single documentation file and extracts metadata.
func parseDocFile(docsDir, filePath string) (*DocFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Parse frontmatter
	var fm FrontMatter
	content, err := frontmatter.Parse(file, &fm)
	if err != nil {
		// If frontmatter parsing fails, try to extract basic info
		fm.Title = extractTitleFromFilename(filePath)
	}

	// Get relative path
	relPath, err := filepath.Rel(docsDir, filePath)
	if err != nil {
		relPath = filePath
	}

	// Use frontmatter title or extract from filename
	title := fm.Title
	if title == "" {
		title = fm.Sidebar
	}
	if title == "" {
		title = extractTitleFromFilename(filePath)
	}

	// Extract first paragraph as description if not in frontmatter
	description := fm.Description
	if description == "" {
		description = extractFirstParagraph(string(content))
	}

	// Infer category and audience from path
	category, audience := inferCategoryAndAudience(relPath)

	return &DocFile{
		Path:        relPath,
		Title:       title,
		Description: truncateString(description, 200),
		Tags:        fm.Tags,
		Category:    category,
		Audience:    audience,
	}, nil
}

// extractTitleFromFilename converts a filename to a title.
// e.g., "feature-flags.mdx" -> "Feature Flags"
func extractTitleFromFilename(path string) string {
	base := filepath.Base(path)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	// Replace hyphens and underscores with spaces
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")

	// Capitalize first letter of each word
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}

	return strings.Join(words, " ")
}

// extractFirstParagraph extracts the first non-empty paragraph from content.
func extractFirstParagraph(content string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	var paragraph strings.Builder
	inParagraph := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines at the beginning
		if !inParagraph && line == "" {
			continue
		}

		// Skip headers, code blocks, and import statements
		if strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "```") ||
			strings.HasPrefix(line, "import ") ||
			strings.HasPrefix(line, ":::") {
			if inParagraph {
				break
			}
			continue
		}

		// Found content
		if line != "" {
			inParagraph = true
			if paragraph.Len() > 0 {
				paragraph.WriteString(" ")
			}
			paragraph.WriteString(line)
		} else if inParagraph {
			// Empty line ends paragraph
			break
		}
	}

	return paragraph.String()
}

// truncateString truncates a string to the specified length, adding "..." if truncated.
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// ReadFile reads the content of a documentation file.
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return string(data), nil
}

// inferCategoryAndAudience derives category and audience from path.
// Based on actual docs structure analysis (78 files in 10 directories).
func inferCategoryAndAudience(path string) (category, audience string) {
	switch {
	// External developers - SDK integration
	case strings.HasPrefix(path, "sdk/"):
		return "sdk", "external-developers"
	case strings.HasPrefix(path, "integration/"):
		return "integration", "external-developers"

	// New users - onboarding
	case strings.HasPrefix(path, "getting-started/"):
		return "getting-started", "new-users"

	// Operators/Product teams - feature management
	case strings.HasPrefix(path, "feature-flags/"):
		return "feature-flags", "operators"
	case strings.HasPrefix(path, "experimentation/"):
		return "experimentation", "operators"

	// Admins - organization management
	case strings.HasPrefix(path, "organization-settings/"):
		return "organization-settings", "admins"

	// All users - changelog (separate sidebar)
	case strings.HasPrefix(path, "changelog/"):
		return "changelog", "all"

	// Engineering leads - best practices
	case strings.HasPrefix(path, "best-practices/"):
		return "best-practices", "engineering-leads"

	// Contributors - contribution guide
	case strings.HasPrefix(path, "contribution-guide/"):
		return "contribution-guide", "contributors"

	default:
		return "general", "all"
	}
}
