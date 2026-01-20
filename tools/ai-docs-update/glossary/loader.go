// Package glossary provides functionality for loading the vocabulary/glossary.
package glossary

import (
	"encoding/json"
	"fmt"
	"os"
)

// Entry represents a single glossary entry.
type Entry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// vocabulary represents the structure of vocabulary.json file.
type vocabulary struct {
	VocabularyList []Entry `json:"vocabularyList"`
}

// Load reads and parses the vocabulary.json file.
// Returns an empty slice if the file doesn't exist or is empty.
func Load(path string) ([]Entry, error) {
	if path == "" {
		return nil, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read glossary file: %w", err)
	}

	if len(data) == 0 {
		return nil, nil
	}

	var vocab vocabulary
	if err := json.Unmarshal(data, &vocab); err != nil {
		return nil, fmt.Errorf("failed to parse glossary JSON: %w", err)
	}

	return vocab.VocabularyList, nil
}

// FormatForPrompt formats glossary entries for inclusion in an AI prompt.
func FormatForPrompt(entries []Entry) string {
	if len(entries) == 0 {
		return ""
	}

	var result string
	for _, entry := range entries {
		result += fmt.Sprintf("- **%s**: %s\n", entry.Name, entry.Description)
	}
	return result
}
