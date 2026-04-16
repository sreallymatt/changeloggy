package hclparse

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/sreallymatt/changeloggy/internal/changes"
)

// EntryFile parses a .hcl changelog entry file and returns the decoded Entries.
func EntryFile(path string) (*changes.Entries, error) {
	entries := &changes.Entries{}
	if err := hclsimple.DecodeFile(path, nil, entries); err != nil {
		return nil, fmt.Errorf("parsing changelog entry file (%s): %w", path, err)
	}
	return entries, nil
}
