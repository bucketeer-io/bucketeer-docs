// Package context provides functionality for loading issue and PR context.
package context

import (
	"fmt"
	"os"
	"strings"
)

// IssueContext holds the context from a GitHub issue.
type IssueContext struct {
	Title string
	Body  string
}

// PRContext holds the context from a GitHub pull request.
type PRContext struct {
	Title string
	Body  string
	Diff  string
}

// LoadIssue loads issue context from the specified files.
// Returns an error if required files cannot be read.
func LoadIssue(titleFile, bodyFile string) (*IssueContext, error) {
	title, err := readFileContent(titleFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read issue title: %w", err)
	}

	body, err := readFileContent(bodyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read issue body: %w", err)
	}

	return &IssueContext{
		Title: strings.TrimSpace(title),
		Body:  strings.TrimSpace(body),
	}, nil
}

// LoadPR loads PR context from the specified files.
// Empty files are allowed (PR may not exist).
func LoadPR(titleFile, bodyFile, diffFile string) (*PRContext, error) {
	ctx := &PRContext{}

	if titleFile != "" {
		title, err := readFileContent(titleFile)
		if err != nil {
			// PR title file is optional; log warning but continue
			title = ""
		}
		ctx.Title = strings.TrimSpace(title)
	}

	if bodyFile != "" {
		body, err := readFileContent(bodyFile)
		if err != nil {
			// PR body file is optional; log warning but continue
			body = ""
		}
		ctx.Body = strings.TrimSpace(body)
	}

	if diffFile != "" {
		diff, err := readFileContent(diffFile)
		if err != nil {
			// Diff file is optional; log warning but continue
			diff = ""
		}
		ctx.Diff = strings.TrimSpace(diff)
	}

	return ctx, nil
}

// String returns a string representation of the issue context.
func (c *IssueContext) String() string {
	return fmt.Sprintf("Issue: %s\n\n%s", c.Title, c.Body)
}

// String returns a string representation of the PR context.
func (c *PRContext) String() string {
	var sb strings.Builder
	if c.Title != "" {
		sb.WriteString(fmt.Sprintf("PR: %s\n\n", c.Title))
	}
	if c.Body != "" {
		sb.WriteString(fmt.Sprintf("Description:\n%s\n\n", c.Body))
	}
	if c.Diff != "" {
		sb.WriteString(fmt.Sprintf("Diff:\n%s", c.Diff))
	}
	return sb.String()
}

// readFileContent reads the entire content of a file.
func readFileContent(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // File doesn't exist, return empty
		}
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return string(data), nil
}
