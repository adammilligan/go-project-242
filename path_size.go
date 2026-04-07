package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize returns the size of the given path as a formatted string.
//
// It supports optional recursion for directories, optional inclusion of hidden entries,
// and optional human-readable output.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, all, recursive)
	if err != nil {
		return "", fmt.Errorf("get path size for %q: %w", path, err)
	}

	return FormatSize(size, human), nil
}

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

// FormatSize formats a size in bytes.
//
// When isHuman is false, it returns bytes (e.g. "123B").
// When isHuman is true, it uses base-1024 units (e.g. "1.2MB").
func FormatSize(size int64, isHuman bool) string {
	const base int64 = 1024

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	if !isHuman {
		return fmt.Sprintf("%d%s", size, units[0])
	}

	if size < base {
		return fmt.Sprintf("%d%s", size, units[0])
	}

	value := float64(size)

	unit := units[0]

	for _, u := range units[1:] {
		if value < float64(base) {
			break
		}

		value /= float64(base)
		unit = u
	}

	return fmt.Sprintf("%.1f%s", value, unit)
}
