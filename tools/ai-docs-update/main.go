// Package main is the entry point for the AI documentation update tool.
// This tool reads issue/PR context and generates documentation updates.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	appctx "github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/context"
	"github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/docs"
	"github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/file"
	"github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/glossary"
	"github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/guardrails"
	"github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/openai"
	"github.com/bucketeer-io/bucketeer-docs/tools/ai-docs-update/styleguide"
)

const appTimeout = 5 * time.Minute

func main() {
	var (
		issueTitleFile = flag.String("issue-title-file", "", "Path to issue title file")
		issueBodyFile  = flag.String("issue-body-file", "", "Path to issue body file")
		prTitleFile    = flag.String("pr-title-file", "", "Path to PR title file")
		prBodyFile     = flag.String("pr-body-file", "", "Path to PR body file")
		diffFile       = flag.String("diff-file", "", "Path to diff file")
		glossaryFile   = flag.String("glossary-file", "", "Path to vocabulary.json file")
		docsDir        = flag.String("docs-dir", "docs", "Path to docs directory")
		excludeDirs    = flag.String("exclude-dirs", "", "Comma-separated directories to exclude (default: changelog,contribution-guide)")
		excludeFiles   = flag.String("exclude-files", "", "Comma-separated files to exclude (default: bucketeer-docs.mdx)")
	)
	flag.Parse()

	// Validate required flags
	if *issueTitleFile == "" || *issueBodyFile == "" {
		log.Fatal("ERROR: --issue-title-file and --issue-body-file are required")
	}

	// Parse exclude lists (nil = use defaults, empty slice = exclude nothing)
	excludeDirsList := parseCommaSeparatedList(*excludeDirs)
	excludeFilesList := parseCommaSeparatedList(*excludeFiles)

	// Set up global timeout
	appCtx, cancel := context.WithTimeout(context.Background(), appTimeout)
	defer cancel()

	if err := run(appCtx, config{
		issueTitleFile: *issueTitleFile,
		issueBodyFile:  *issueBodyFile,
		prTitleFile:    *prTitleFile,
		prBodyFile:     *prBodyFile,
		diffFile:       *diffFile,
		glossaryFile:   *glossaryFile,
		docsDir:        *docsDir,
		excludeDirs:    excludeDirsList,
		excludeFiles:   excludeFilesList,
	}); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

type config struct {
	issueTitleFile string
	issueBodyFile  string
	prTitleFile    string
	prBodyFile     string
	diffFile       string
	glossaryFile   string
	docsDir        string
	excludeDirs    []string // nil = use defaults, empty slice = exclude nothing
	excludeFiles   []string // nil = use defaults, empty slice = exclude nothing
}

func run(ctx context.Context, cfg config) error {
	// 1. Load Issue context
	issueCtx, err := appctx.LoadIssue(cfg.issueTitleFile, cfg.issueBodyFile)
	if err != nil {
		return fmt.Errorf("failed to load issue context: %w", err)
	}
	log.Printf("Issue: %s", issueCtx.Title)

	// 2. Load PR context
	prCtx, err := appctx.LoadPR(cfg.prTitleFile, cfg.prBodyFile, cfg.diffFile)
	if err != nil {
		return fmt.Errorf("failed to load PR context: %w", err)
	}
	if prCtx.Title != "" {
		log.Printf("PR: %s", prCtx.Title)
	} else {
		log.Println("PR: (no PR context provided)")
	}

	// 3. Load glossary (optional - continue without it if loading fails)
	glossaryEntries, err := glossary.Load(cfg.glossaryFile)
	if err != nil {
		log.Printf("Warning: Failed to load glossary: %v (continuing without glossary)", err)
		glossaryEntries = nil
	} else {
		log.Printf("Loaded %d glossary entries", len(glossaryEntries))
	}

	// 3.5. Load style guide (optional - continue with defaults if loading fails)
	styleGuideDir := filepath.Join(cfg.docsDir, "contribution-guide", "documentation-style")
	styleGuideData, err := styleguide.Load(styleGuideDir)
	if err != nil {
		log.Printf("Warning: Failed to load style guide: %v (using defaults)", err)
	} else {
		log.Printf("Loaded %d style guide rules", styleGuideData.RuleCount())
	}
	formattedStyleGuide := styleGuideData.Format()

	// 4. Input guardrails validation
	inputGuard := guardrails.NewInputGuardrails()
	if err := inputGuard.ValidateContext(toGuardrailsIssueContext(issueCtx), toGuardrailsPRContext(prCtx)); err != nil {
		log.Printf("Input guardrails triggered (skipping): %v", err)
		return nil // Skip without error - this is expected behavior
	}

	// 4.5. Validate diff structure (file count, line count)
	if prCtx.Diff != "" {
		parsedDiff := guardrails.ParseDiff(prCtx.Diff)
		if err := inputGuard.ValidateDiff(parsedDiff); err != nil {
			log.Printf("Diff guardrails triggered (skipping): %v", err)
			return nil
		}
		log.Printf("Diff validated: %d files, %d bytes", len(parsedDiff.Files), parsedDiff.TotalSize)
	}

	// 5. Generate docs manifest (nil = use defaults for exclusions)
	manifest, err := docs.GenerateManifest(cfg.docsDir, cfg.excludeDirs, cfg.excludeFiles)
	if err != nil {
		return fmt.Errorf("failed to generate docs manifest: %w", err)
	}
	log.Printf("Found %d documentation files (excluded dirs: %v, excluded files: %v)",
		len(manifest.Files),
		withDefault(cfg.excludeDirs, docs.DefaultExcludeDirs),
		withDefault(cfg.excludeFiles, docs.DefaultExcludeFiles))

	// 6. Log context summary (for debugging)
	logContextSummary(issueCtx, prCtx, glossaryEntries, manifest)

	// Check for context cancellation
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancelled: %w", ctx.Err())
	default:
	}

	// 7. Initialize OpenAI client
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}
	client := openai.NewClient(apiKey)

	// 8. Phase 1: AI identifies which docs to update
	log.Println("Phase 1: Identifying documents to update...")

	// Create diff summary for Phase 1 (efficient format, not full diff)
	diffSummary := guardrails.SummarizeDiff(prCtx.Diff)

	// Token limit check for Phase 1 (includes glossary, manifest, diff summary)
	phase1TokenEstimate := estimatePhase1Tokens(issueCtx, prCtx, glossaryEntries, manifest, diffSummary)
	if phase1TokenEstimate > guardrails.MaxInputTokens {
		log.Printf("Phase 1 token limit exceeded: ~%d tokens (max %d)", phase1TokenEstimate, guardrails.MaxInputTokens)
		return fmt.Errorf("phase 1 context too large: ~%d tokens", phase1TokenEstimate)
	}
	log.Printf("Phase 1 token estimate: ~%d tokens", phase1TokenEstimate)

	identification, err := client.IdentifyDocsToUpdate(ctx, openai.IdentifyRequest{
		IssueTitle:   issueCtx.Title,
		IssueBody:    issueCtx.Body,
		PRTitle:      prCtx.Title,
		PRBody:       prCtx.Body,
		DiffSummary:  diffSummary,
		Glossary:     toOpenAIGlossary(glossaryEntries),
		DocsManifest: toOpenAIDocsManifest(manifest),
	})
	if err != nil {
		return fmt.Errorf("failed to identify docs: %w", err)
	}

	if !identification.NeedsUpdate {
		log.Printf("AI determined no docs need updating: %s", identification.Reason)
		return nil
	}

	log.Printf("AI identified %d files to update", len(identification.FilesToUpdate))
	for _, f := range identification.FilesToUpdate {
		log.Printf("  - %s (%s): %s", f.Path, f.UpdateType, f.BriefDescription)
	}

	// 9. Phase 2: Generate updates for each identified file
	log.Println("Phase 2: Generating document updates...")
	outputGuard := guardrails.NewOutputGuardrails()
	var successCount int

	// Create Writer with manifest paths for validation
	manifestPaths := getManifestPaths(manifest)
	writer := file.NewWriter(cfg.docsDir, manifestPaths)

	for _, fileUpdate := range identification.FilesToUpdate {
		log.Printf("Processing: %s (%s)", fileUpdate.Path, fileUpdate.UpdateType)

		// Build full path for reading (docsDir + relative path from manifest)
		fullPath := filepath.Join(cfg.docsDir, fileUpdate.Path)

		// Look up content type from manifest
		contentType := findContentType(manifest, fileUpdate.Path)

		// Read current content
		currentContent, err := docs.ReadFile(fullPath)
		if err != nil {
			log.Printf("ERROR: Failed to read %s: %v (skipping)", fileUpdate.Path, err)
			continue
		}

		// Validate document content size
		if err := inputGuard.ValidateDocContent(currentContent); err != nil {
			log.Printf("Document too large for %s (skipping): %v", fileUpdate.Path, err)
			continue
		}

		// Token limit check
		combinedContext := issueCtx.String() + prCtx.String()
		if err := inputGuard.ValidateTokenLimit(combinedContext, currentContent); err != nil {
			log.Printf("Token limit exceeded for %s (skipping): %v", fileUpdate.Path, err)
			continue
		}

		// Generate update (Full file content)
		rawResult, err := client.GenerateDocUpdate(ctx, openai.UpdateRequest{
			IssueTitle:        issueCtx.Title,
			IssueBody:         issueCtx.Body,
			PRTitle:           prCtx.Title,
			PRBody:            prCtx.Body,
			CodeDiff:          prCtx.Diff,
			Glossary:          toOpenAIGlossary(glossaryEntries),
			DocPath:           fileUpdate.Path,
			CurrentContent:    currentContent,
			UpdateInstruction: fileUpdate.BriefDescription,
			ContentType:       contentType,
			StyleGuide:        formattedStyleGuide,
		})
		if err != nil {
			log.Printf("ERROR: OpenAI error for %s: %v (skipping)", fileUpdate.Path, err)
			continue
		}

		// Output guardrails validation
		if err := outputGuard.Validate(rawResult); err != nil {
			log.Printf("Output guardrails triggered for %s: %v (skipping)", fileUpdate.Path, err)
			continue
		}

		// Extract document content
		content := guardrails.ExtractDocumentContent(rawResult)
		if content == "" {
			log.Printf("ERROR: Failed to extract content for %s (skipping)", fileUpdate.Path)
			continue
		}

		// Apply post-processing transformations (placeholder→TODO, non-ASCII→ASCII)
		content, postProcessWarnings := outputGuard.PostProcess(content)
		for _, w := range postProcessWarnings {
			log.Printf("Post-process warning for %s: %s", fileUpdate.Path, w)
		}

		// Log TODO markers if any
		if openai.HasTODOMarkers(content) {
			log.Printf("TODO markers found in %s (requires manual review)", fileUpdate.Path)
		}

		// Write file (validates path is in manifest)
		if err := writer.Write(fileUpdate.Path, content); err != nil {
			log.Printf("ERROR: Failed to write %s: %v (skipping)", fileUpdate.Path, err)
			continue
		}
		log.Printf("Successfully updated: %s", fileUpdate.Path)

		successCount++
	}

	log.Printf("Completed: %d/%d files updated", successCount, len(identification.FilesToUpdate))
	return nil
}

func logContextSummary(
	issueCtx *appctx.IssueContext,
	prCtx *appctx.PRContext,
	glossaryEntries []glossary.Entry,
	manifest *docs.Manifest,
) {
	log.Println("=== Context Summary ===")
	log.Printf("Issue Title: %s", issueCtx.Title)
	log.Printf("Issue Body Length: %d bytes", len(issueCtx.Body))

	if prCtx.Title != "" {
		log.Printf("PR Title: %s", prCtx.Title)
		log.Printf("PR Body Length: %d bytes", len(prCtx.Body))
		log.Printf("Diff Length: %d bytes", len(prCtx.Diff))
	}

	if glossaryEntries != nil {
		log.Printf("Glossary Entries: %d", len(glossaryEntries))
	}

	log.Printf("Documentation Files: %d", len(manifest.Files))
	log.Println("=======================")
}

// Type conversion helpers

// toGuardrailsIssueContext converts context.IssueContext to guardrails.IssueContext
func toGuardrailsIssueContext(ic *appctx.IssueContext) *guardrails.IssueContext {
	if ic == nil {
		return nil
	}
	return &guardrails.IssueContext{
		Title: ic.Title,
		Body:  ic.Body,
	}
}

// toGuardrailsPRContext converts context.PRContext to guardrails.PRContext
func toGuardrailsPRContext(pc *appctx.PRContext) *guardrails.PRContext {
	if pc == nil {
		return nil
	}
	return &guardrails.PRContext{
		Title: pc.Title,
		Body:  pc.Body,
		Diff:  pc.Diff,
	}
}

// toOpenAIGlossary converts glossary.Entry slice to openai.GlossaryEntry slice
func toOpenAIGlossary(entries []glossary.Entry) []openai.GlossaryEntry {
	if entries == nil {
		return nil
	}
	result := make([]openai.GlossaryEntry, len(entries))
	for i, e := range entries {
		result[i] = openai.GlossaryEntry{
			Name:        e.Name,
			Description: e.Description,
		}
	}
	return result
}

// toOpenAIDocsManifest converts docs.Manifest to openai.DocsManifest
func toOpenAIDocsManifest(m *docs.Manifest) *openai.DocsManifest {
	if m == nil {
		return nil
	}
	files := make([]openai.DocFile, len(m.Files))
	for i, f := range m.Files {
		files[i] = openai.DocFile{
			Path:        f.Path,
			Title:       f.Title,
			Description: f.Description,
			Tags:        f.Tags,
			Category:    f.Category,
			Audience:    f.Audience,
			ContentType: string(f.ContentType),
		}
	}
	return &openai.DocsManifest{Files: files}
}

// parseCommaSeparatedList splits a comma-separated string into a trimmed slice.
// Returns nil for empty input (to use defaults), or a slice of trimmed values.
func parseCommaSeparatedList(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// findContentType looks up the content type for a file path from the manifest.
func findContentType(m *docs.Manifest, path string) string {
	for _, f := range m.Files {
		if f.Path == path {
			return string(f.ContentType)
		}
	}
	return "user-guide"
}

// withDefault returns the slice if non-nil, otherwise returns the default.
func withDefault(slice, defaultSlice []string) []string {
	if slice == nil {
		return defaultSlice
	}
	return slice
}

// getManifestPaths extracts all file paths from the manifest.
func getManifestPaths(m *docs.Manifest) []string {
	if m == nil {
		return nil
	}
	paths := make([]string, len(m.Files))
	for i, f := range m.Files {
		paths[i] = f.Path
	}
	return paths
}

// estimatePhase1Tokens estimates token count for Phase 1 prompt.
// Includes: issue, PR, glossary, manifest, diff summary, and prompt template overhead.
func estimatePhase1Tokens(
	issue *appctx.IssueContext,
	pr *appctx.PRContext,
	glossaryEntries []glossary.Entry,
	manifest *docs.Manifest,
	diffSummary string,
) int {
	// Base overhead for prompt template (~500 tokens)
	tokens := 500

	// Issue context
	if issue != nil {
		tokens += estimateTokens(issue.Title)
		tokens += estimateTokens(issue.Body)
	}

	// PR context
	if pr != nil {
		tokens += estimateTokens(pr.Title)
		tokens += estimateTokens(pr.Body)
	}

	// Diff summary
	tokens += estimateTokens(diffSummary)

	// Glossary (estimate ~20 tokens per entry)
	tokens += len(glossaryEntries) * 20

	// Manifest (estimate ~30 tokens per file entry)
	if manifest != nil {
		tokens += len(manifest.Files) * 30
	}

	return tokens
}

// estimateTokens provides simple token estimation (1 token ≈ 4 chars).
func estimateTokens(text string) int {
	return len(text) / 4
}
