package sizecalc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetSize calculates the total size in bytes for the given path.
//
// If path is a file, it returns the file size. If path is a directory, it sums the sizes of
// the directory entries. When recursive is true, nested directories are included as well.
// When all is false, entries with names starting with '.' are skipped.
func GetSize(path string, all bool, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("stat %q: %w", path, err)
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("read dir %q: %w", path, err)
	}

	var total int64

	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("stat entry %q in %q: %w", entry.Name(), path, err)
		}

		if entryInfo.IsDir() {
			if !recursive {
				continue
			}

			subPath := filepath.Join(path, entry.Name())

			subSize, err := GetSize(subPath, all, recursive)
			if err != nil {
				return 0, fmt.Errorf("get size for %q: %w", subPath, err)
			}

			total += subSize

			continue
		}

		total += entryInfo.Size()
	}

	return total, nil
}
