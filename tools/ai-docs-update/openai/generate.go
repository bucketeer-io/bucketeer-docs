package openai

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"
)

// DocUpdatePrompt is the prompt template for Phase 2: Update Generation.
// It generates the updated documentation content based on issue/PR context.
const DocUpdatePrompt = `You are a technical documentation updater for Bucketeer,
a feature flag and A/B testing platform.

## GLOSSARY (Use these terms consistently)
{{range .Glossary}}
- **{{.Name}}**: {{.Description}}
{{end}}

## ISSUE CONTEXT
<issue>
<title>{{.SanitizedIssueTitle}}</title>
<body>{{.SanitizedIssueBody}}</body>
</issue>

## PR CONTEXT
<pr_context>
<title>{{.SanitizedPRTitle}}</title>
<description>{{.SanitizedPRBody}}</description>
</pr_context>

## CODE CHANGES (for technical reference)
<code_diff>
{{.SanitizedDiff}}
</code_diff>

## DOCUMENT TO UPDATE
File: {{.DocPath}}
<current_document>
{{.CurrentContent}}
</current_document>

## UPDATE INSTRUCTION
{{.UpdateInstruction}}

## RULES (MUST FOLLOW)
1. Use Issue body and PR description as the PRIMARY source of truth for feature explanation
2. Use code diff only for technical accuracy (field names, API endpoints, etc.)
3. Do NOT invent information not in issue/PR description or code
4. If information is unclear, add "TODO: Needs confirmation" instead of guessing
5. Maintain the existing document structure and style
6. Keep changes minimal and focused
7. Preserve all existing content that is not being updated
8. Use terminology from the GLOSSARY consistently
9. **NEVER use version placeholders** like "X.Y.Z", "vN.N.N", "TBD", or similar. Use "TODO: Needs confirmation - version number" instead
10. If a specific version number is not provided in the issue/PR, omit version references entirely or use TODO markers

## STYLE GUIDE
- Use contractions (don't, it's, you'll) for a friendly tone
- Use backticks for code elements: field names, API endpoints, values
- Use present tense for feature descriptions
- Be concise and direct

## CHARACTER ENCODING (MUST FOLLOW)
- Use only ASCII characters for punctuation:
  - Hyphen: - (U+002D), NOT en-dash (U+2013) or em-dash (U+2014)
  - Apostrophe: ' (U+0027), NOT curly quotes (U+2019)
  - Quotation marks: " (U+0022), NOT curly quotes (U+201C, U+201D)
  - Ellipsis: ... (three periods), NOT ellipsis character (U+2026)
- Exception: Curly quotes are allowed in quoted content from external sources

## OUTPUT FORMAT
Return the COMPLETE updated document content wrapped in tags.
Do NOT output a diff - output the full file content with your changes applied.

<updated_document>
[Full document content here]
</updated_document>
`

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
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// NOTE: ExtractDocumentContent has been consolidated into guardrails package.
// Use guardrails.ExtractDocumentContent() for extracting document content.

// HasDocumentTags checks if the output contains the required document tags.
func HasDocumentTags(output string) bool {
	return strings.Contains(output, "<updated_document>") &&
		strings.Contains(output, "</updated_document>")
}

// HasTODOMarkers checks if the content contains TODO markers that need review.
func HasTODOMarkers(content string) bool {
	return strings.Contains(content, "TODO: Needs confirmation")
}
