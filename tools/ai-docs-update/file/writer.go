// Package file provides functionality for writing documentation files.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ErrPathTraversal indicates an attempt to write outside allowed directories
var ErrPathTraversal = errors.New("path traversal detected")

// Write writes content to a file, creating directories if needed.
// Uses atomic write pattern: write to temp file, then rename.
func Write(path string, content string) error {
	// Path traversal protection
	if err := validatePath(path); err != nil {
		return err
	}

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Create temp file in the same directory for atomic rename
	tempFile, err := os.CreateTemp(dir, ".ai-docs-update-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tempPath := tempFile.Name()

	// Clean up temp file on error
	defer func() {
		if tempPath != "" {
			os.Remove(tempPath)
		}
	}()

	// Write content to temp file
	if _, err := tempFile.WriteString(content); err != nil {
		tempFile.Close()
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	// Sync to ensure data is flushed to disk
	if err := tempFile.Sync(); err != nil {
		tempFile.Close()
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tempPath, path); err != nil {
		return fmt.Errorf("failed to rename temp file to %s: %w", path, err)
	}

	// Clear tempPath so defer doesn't try to remove it
	tempPath = ""

	return nil
}

// WriteWithBackup writes content to a file, creating a backup of the original.
func WriteWithBackup(path string, content string) error {
	// Check if original file exists
	if _, err := os.Stat(path); err == nil {
		// Create backup
		backupPath := path + ".bak"
		if err := copyFile(path, backupPath); err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}
	}

	return Write(path, content)
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

// validatePath checks for path traversal attempts
func validatePath(path string) error {
	// Clean the path first
	cleanPath := filepath.Clean(path)

	// Ensure path contains "docs" directory (basic safety check)
	// This allows paths like "../../docs/feature-flags/segments.mdx"
	if !strings.Contains(cleanPath, "docs"+string(filepath.Separator)) &&
		!strings.Contains(cleanPath, "docs/") &&
		!strings.HasPrefix(cleanPath, "docs") {
		return fmt.Errorf("%w: path not within docs directory: %s", ErrPathTraversal, path)
	}

	// Check that the path doesn't try to escape after docs/
	// e.g., "docs/../../../etc/passwd" should be rejected
	parts := strings.Split(cleanPath, string(filepath.Separator))
	docsFound := false
	depthAfterDocs := 0
	for _, part := range parts {
		if part == "docs" {
			docsFound = true
			depthAfterDocs = 0
			continue
		}
		if docsFound {
			if part == ".." {
				depthAfterDocs--
				if depthAfterDocs < 0 {
					return fmt.Errorf("%w: path escapes docs directory: %s", ErrPathTraversal, path)
				}
			} else if part != "." && part != "" {
				depthAfterDocs++
			}
		}
	}

	return nil
}
