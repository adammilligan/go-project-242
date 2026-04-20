package code

import (
	"fmt"

	"code/internal/format"
	"code/internal/sizecalc"
)

// GetPathSize returns the size of the given path as a formatted string.
//
// It supports optional recursion for directories, optional inclusion of hidden entries,
// and optional human-readable output.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := sizecalc.GetSize(path, all, recursive)
	if err != nil {
		return "", fmt.Errorf("get path size for %q: %w", path, err)
	}

	return format.FormatSize(size, human), nil
}
