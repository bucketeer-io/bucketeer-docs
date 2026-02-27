package openai

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

// DocIdentificationPrompt is the prompt template for Phase 1: Document Identification.
// It analyzes issue/PR context to determine which documentation files need updates.
//
//go:embed prompts/identify.tmpl
var DocIdentificationPrompt string

// GlossaryEntry represents a term in the glossary.
type GlossaryEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DocFile represents a documentation file in the manifest.
type DocFile struct {
	Path        string   `json:"path"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Category    string   `json:"category,omitempty"`
	Audience    string   `json:"audience,omitempty"`
	ContentType string   `json:"content_type,omitempty"` // user-guide, admin-config, developer-reference
}

// DocsManifest represents the list of available documentation files.
type DocsManifest struct {
	Files []DocFile `json:"files"`
}

// IdentifyRequest contains all data needed for document identification.
type IdentifyRequest struct {
	IssueTitle   string
	IssueBody    string
	PRTitle      string
	PRBody       string
	DiffSummary  string // Summary of changed files (not full diff)
	Glossary     []GlossaryEntry
	DocsManifest *DocsManifest
}

// FileToUpdate represents a file that needs to be updated.
type FileToUpdate struct {
	Path             string `json:"path"`
	UpdateType       string `json:"update_type"`
	BriefDescription string `json:"brief_description"`
	TargetLocation   string `json:"target_location"` // Where to add content (for add_inline/modify_section)
}

// IdentifyResponse represents the AI's response for document identification.
type IdentifyResponse struct {
	NeedsUpdate   bool           `json:"needs_update"`
	Reason        string         `json:"reason"`
	FilesToUpdate []FileToUpdate `json:"files_to_update"`
}

// identifyTemplateData is the data structure for the identification prompt template.
type identifyTemplateData struct {
	Glossary      []GlossaryEntry
	IssueTitle    string
	IssueBody     string
	PRTitle       string
	PRDescription string
	DiffSummary   string
	DocsManifest  *DocsManifest
}

// IdentifyDocsToUpdate executes Phase 1: Document Identification.
// It analyzes the issue/PR context and returns which documentation files need updates.
func (c *Client) IdentifyDocsToUpdate(ctx context.Context, req IdentifyRequest) (*IdentifyResponse, error) {
	// Build prompt from template
	prompt, err := buildIdentifyPrompt(req)
	if err != nil {
		return nil, fmt.Errorf("failed to build identify prompt: %w", err)
	}

	// Create messages
	messages := []ChatMessage{
		{
			Role:    "system",
			Content: "You are a documentation analyst. Respond only with valid JSON.",
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	// Call OpenAI API with JSON response format to guarantee valid JSON output
	response, err := c.CreateChatCompletion(ctx, messages, WithJSONResponse())
	if err != nil {
		return nil, fmt.Errorf("openai api call failed: %w", err)
	}

	// Parse response
	result, err := parseIdentifyResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse identify response: %w", err)
	}

	return result, nil
}

// buildIdentifyPrompt builds the prompt for document identification.
func buildIdentifyPrompt(req IdentifyRequest) (string, error) {
	tmpl, err := template.New("identify").Parse(DocIdentificationPrompt)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Sanitize inputs
	sanitized := NewSanitizedContext(
		req.IssueTitle,
		req.IssueBody,
		req.PRTitle,
		req.PRBody,
		"", // Full diff not needed for identification - we use summary
	)

	data := identifyTemplateData{
		Glossary:      req.Glossary,
		IssueTitle:    sanitized.IssueTitle,
		IssueBody:     sanitized.IssueBody,
		PRTitle:       sanitized.PRTitle,
		PRDescription: sanitized.PRBody,
		DiffSummary:   req.DiffSummary,
		DocsManifest:  req.DocsManifest,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// parseIdentifyResponse parses the AI response into IdentifyResponse.
func parseIdentifyResponse(response string) (*IdentifyResponse, error) {
	// Try to extract JSON from the response
	// The response might contain markdown code blocks
	jsonStr := extractJSON(response)

	var result IdentifyResponse
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w (response: %s)", err, response)
	}

	// Validate the response
	if result.NeedsUpdate && len(result.FilesToUpdate) == 0 {
		return nil, fmt.Errorf("invalid response: needs_update is true but no files_to_update")
	}

	// Limit to maximum 3 files
	if len(result.FilesToUpdate) > 3 {
		result.FilesToUpdate = result.FilesToUpdate[:3]
	}

	return &result, nil
}

// extractJSON attempts to extract JSON from a response that might contain markdown.
// This serves as a defensive fallback even when response_format: json_object is enabled,
// because the JSON guarantee may not hold for all models or proxies configurable
// via OPENAI_MODEL / OPENAI_API_BASE.
func extractJSON(response string) string {
	// Look for ```json or ``` block
	const jsonBlockStart = "```json"
	const codeBlockMarker = "```"

	start := -1
	startIdx := strings.Index(response, jsonBlockStart)
	if startIdx != -1 {
		start = startIdx + len(jsonBlockStart)
	} else {
		// Try plain ``` block
		startIdx = strings.Index(response, codeBlockMarker)
		if startIdx != -1 {
			start = startIdx + len(codeBlockMarker)
		}
	}

	if start != -1 {
		// Find the closing ```
		remaining := response[start:]
		endIdx := strings.Index(remaining, codeBlockMarker)
		if endIdx != -1 {
			return response[start : start+endIdx]
		}
	}

	// Look for raw JSON (starts with { and ends with })
	braceStart := strings.Index(response, "{")
	braceEnd := strings.LastIndex(response, "}")
	if braceStart != -1 && braceEnd != -1 && braceEnd > braceStart {
		return response[braceStart : braceEnd+1]
	}

	// Return as-is if no JSON found
	return response
}
