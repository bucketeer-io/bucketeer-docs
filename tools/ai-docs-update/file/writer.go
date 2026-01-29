// Package file provides functionality for writing documentation files.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ErrPathNotInManifest indicates the path is not in the allowed manifest
var ErrPathNotInManifest = errors.New("path not in manifest")

// Writer provides file writing with manifest validation.
type Writer struct {
	rootDir      string
	allowedPaths map[string]bool
}

// NewWriter creates a Writer that only allows writing to manifest paths.
func NewWriter(rootDir string, manifestPaths []string) *Writer {
	allowed := make(map[string]bool, len(manifestPaths))
	for _, p := range manifestPaths {
		allowed[filepath.ToSlash(p)] = true
	}
	return &Writer{
		rootDir:      rootDir,
		allowedPaths: allowed,
	}
}

// Write writes content to a file if the path is in the manifest.
func (w *Writer) Write(relativePath, content string) error {
	// Validate path is in manifest
	normalized := filepath.ToSlash(relativePath)
	if !w.allowedPaths[normalized] {
		return fmt.Errorf("%w: %s", ErrPathNotInManifest, relativePath)
	}

	fullPath := filepath.Join(w.rootDir, relativePath)

	// Ensure directory exists
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write atomically via temp file
	tempFile, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	if _, err := tempFile.WriteString(content); err != nil {
		tempFile.Close()
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := tempFile.Sync(); err != nil {
		tempFile.Close()
		return fmt.Errorf("failed to sync: %w", err)
	}
	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed to close: %w", err)
	}

	if err := os.Rename(tempPath, fullPath); err != nil {
		return fmt.Errorf("failed to rename: %w", err)
	}

	return nil
}
