package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EnsureDir(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("creating directory (%s): %w", dir, err)
	}
	return nil
}

// Code wraps a value in backticks.
// If the value contains backticks, it adds additional ones to ensure it renders properly as a code block.
// If the value begins or ends with backticks, it adds the requred spacing ensure it renders properly as a code block.
func Code(value string) string {
	delimiter := "`"
	for strings.Contains(value, delimiter) {
		delimiter += "`"
	}

	if strings.HasPrefix(value, "`") || strings.HasSuffix(value, "`") {
		value = " " + value + " "
	}

	return delimiter + value + delimiter
}
