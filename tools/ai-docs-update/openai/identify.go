package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"
)

// DocIdentificationPrompt is the prompt template for Phase 1: Document Identification.
// It analyzes issue/PR context to determine which documentation files need updates.
const DocIdentificationPrompt = `You are a documentation analyst for Bucketeer,
a feature flag and A/B testing platform.

## GLOSSARY (Use these terms consistently)
{{range .Glossary}}
- **{{.Name}}**: {{.Description}}
{{end}}

## CRITICAL: AUDIENCE DISTINCTION
The Bucketeer project has TWO distinct codebases:

1. **Bucketeer SDKs** (for EXTERNAL developers)
   - npm packages: @bucketeer/js-client-sdk, @bucketeer/react-client-sdk
   - Native SDKs: Android, iOS, Flutter
   - Server SDKs: Go, Node.js
   - Documentation in: /docs/sdk/**
   - Audience: External developers integrating feature flags into THEIR applications

2. **Bucketeer Dashboard** (for INTERNAL operators)
   - React web application in ui/dashboard/src/
   - Used by Bucketeer operators to manage feature flags
   - Documentation in: /docs/getting-started/bucketeer-dashboard.mdx and dashboard-related guides
   - Audience: Operators using the Bucketeer admin console

**NEVER** document internal Dashboard code (React hooks, components from ui/dashboard/src/) in SDK documentation.

## TASK
Analyze the following feature change and identify which documentation files need to be updated.

## ISSUE CONTEXT
Issue Title: {{.IssueTitle}}
Issue Body:
{{.IssueBody}}

## LINKED PR
PR Title: {{.PRTitle}}
PR Description:
{{.PRDescription}}

## AVAILABLE DOCUMENTATION FILES
{{range .DocsManifest.Files}}
- {{.Path}} [{{.Category}}|{{.Audience}}|{{.ContentType}}]: {{.Title}}
{{end}}

## CONTENT TYPE DEFINITIONS
- **user-guide**: User-facing behavior docs (what users see/experience). NO implementation details.
- **admin-config**: Configuration docs for operators (CLI flags, Helm values, env vars).
- **developer-reference**: SDK/API reference for external developers (public methods, integration code).

## OUTPUT FORMAT (JSON only)
{
  "needs_update": true/false,
  "reason": "brief explanation",
  "files_to_update": [
    {
      "path": "feature-flags/xxx.mdx",
      "update_type": "add_section|modify_section|add_example",
      "brief_description": "what to add/change"
    }
  ]
}

## RULES
1. Only select files that are DIRECTLY related to the feature
2. If the feature is entirely new and no existing doc covers it, set needs_update to false and explain
3. Prefer updating existing sections over creating new ones
4. Maximum 3 files per feature change
5. **CRITICAL**: Match audience - SDK changes go to SDK docs, Dashboard changes go to dashboard docs
6. If the PR modifies ui/dashboard/src/**, do NOT update /docs/sdk/** files
7. If the PR modifies SDK packages (@bucketeer/*-sdk), do NOT update dashboard operation guides

## SINGLE SOURCE OF TRUTH (CRITICAL - Prevents Duplication)
8. **Each piece of information should appear in ONLY ONE document.**
   - Per-environment configuration → environments.mdx (NOT settings.mdx)
   - Per-organization configuration → settings.mdx
   - User-facing dashboard behavior → bucketeer-dashboard.mdx
   - SDK integration details → sdk/**
9. **When information could fit multiple files, choose the MOST SPECIFIC one.**
   - Token TTL config (per-env) → environments.mdx only
   - Organization name/URL → settings.mdx only
10. **Cross-reference instead of duplicate.** If a doc needs to mention related content, link to the authoritative doc instead of repeating the information.
`

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
	Glossary     []GlossaryEntry
	DocsManifest *DocsManifest
}

// FileToUpdate represents a file that needs to be updated.
type FileToUpdate struct {
	Path             string `json:"path"`
	UpdateType       string `json:"update_type"`
	BriefDescription string `json:"brief_description"`
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

	// Call OpenAI API
	response, err := c.CreateChatCompletion(ctx, messages)
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
		"", // No diff needed for identification
	)

	data := identifyTemplateData{
		Glossary:      req.Glossary,
		IssueTitle:    sanitized.IssueTitle,
		IssueBody:     sanitized.IssueBody,
		PRTitle:       sanitized.PRTitle,
		PRDescription: sanitized.PRBody,
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
func extractJSON(response string) string {
	// Try to find JSON in code blocks
	start := -1
	end := -1

	// Look for ```json or ``` block
	jsonBlockStart := "```json"
	jsonBlockEnd := "```"

	startIdx := bytes.Index([]byte(response), []byte(jsonBlockStart))
	if startIdx != -1 {
		start = startIdx + len(jsonBlockStart)
	} else {
		// Try plain ``` block
		startIdx = bytes.Index([]byte(response), []byte("```"))
		if startIdx != -1 {
			start = startIdx + len("```")
		}
	}

	if start != -1 {
		// Find the closing ```
		remaining := response[start:]
		endIdx := bytes.Index([]byte(remaining), []byte(jsonBlockEnd))
		if endIdx != -1 {
			end = start + endIdx
		}
	}

	if start != -1 && end != -1 {
		return response[start:end]
	}

	// Look for raw JSON (starts with { and ends with })
	braceStart := bytes.Index([]byte(response), []byte("{"))
	braceEnd := bytes.LastIndex([]byte(response), []byte("}"))
	if braceStart != -1 && braceEnd != -1 && braceEnd > braceStart {
		return response[braceStart : braceEnd+1]
	}

	// Return as-is if no JSON found
	return response
}
