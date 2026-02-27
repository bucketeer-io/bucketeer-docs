package openai

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
)

// DocUpdatePrompt is the prompt template for Phase 2: Update Generation.
// It generates the updated documentation content based on issue/PR context.
//
//go:embed prompts/update.tmpl
var DocUpdatePrompt string

// UpdateRequest contains all data needed for document update generation.
type UpdateRequest struct {
	IssueTitle        string
	IssueBody         string
	PRTitle           string
	PRBody            string
	CodeDiff          string
	Glossary          []GlossaryEntry
	DocPath           string
	CurrentContent    string
	UpdateInstruction string
	ContentType       string // user-guide, admin-config, developer-reference
	StyleGuide        string // Formatted style guide rules
	UpdateType        string // add_inline, modify_section, add_section, add_example
}

// updateTemplateData is the data structure for the update prompt template.
type updateTemplateData struct {
	Glossary            []GlossaryEntry
	SanitizedIssueTitle string
	SanitizedIssueBody  string
	SanitizedPRTitle    string
	SanitizedPRBody     string
	SanitizedDiff       string
	DocPath             string
	CurrentContent      string
	UpdateInstruction   string
	ContentType         string
	StyleGuide          string
	UpdateType          string
}

// GenerateDocUpdate executes Phase 2: Update Generation.
// It generates the updated documentation content based on the context provided.
func (c *Client) GenerateDocUpdate(ctx context.Context, req UpdateRequest) (string, error) {
	// Build prompt from template
	prompt, err := buildUpdatePrompt(req)
	if err != nil {
		return "", fmt.Errorf("failed to build update prompt: %w", err)
	}

	// Create messages
	messages := []ChatMessage{
		{
			Role:    "system",
			Content: "You are a technical documentation updater. Follow the rules exactly and output the complete updated document wrapped in <updated_document> tags.",
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	// Call OpenAI API
	response, err := c.CreateChatCompletion(ctx, messages)
	if err != nil {
		return "", fmt.Errorf("openai api call failed: %w", err)
	}

	return response, nil
}

// buildUpdatePrompt builds the prompt for document update generation.
func buildUpdatePrompt(req UpdateRequest) (string, error) {
	tmpl, err := template.New("update").Parse(DocUpdatePrompt)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Sanitize inputs
	sanitized := NewSanitizedContext(
		req.IssueTitle,
		req.IssueBody,
		req.PRTitle,
		req.PRBody,
		req.CodeDiff,
	)

	data := updateTemplateData{
		Glossary:            req.Glossary,
		SanitizedIssueTitle: sanitized.IssueTitle,
		SanitizedIssueBody:  sanitized.IssueBody,
		SanitizedPRTitle:    sanitized.PRTitle,
		SanitizedPRBody:     sanitized.PRBody,
		SanitizedDiff:       sanitized.Diff,
		DocPath:             req.DocPath,
		CurrentContent:      req.CurrentContent,
		UpdateInstruction:   req.UpdateInstruction,
		ContentType:         req.ContentType,
		StyleGuide:          req.StyleGuide,
		UpdateType:          req.UpdateType,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// HasDocumentTags checks if the output contains the required document tags.
func HasDocumentTags(output string) bool {
	return strings.Contains(output, "<updated_document>") &&
		strings.Contains(output, "</updated_document>")
}
