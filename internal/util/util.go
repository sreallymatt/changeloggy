package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureDir(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating directory (%s): %w", dir, err)
	}
	return nil
}
