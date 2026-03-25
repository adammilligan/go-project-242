package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, all, recursive)
	if err != nil {
		return "", err
	}

	return FormatSize(size, human), nil
}

func GetSize(path string, all bool, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		if entryInfo.IsDir() {
			if !recursive {
				continue
			}

			subPath := filepath.Join(path, entry.Name())
			subSize, err := GetSize(subPath, all, recursive)
			if err != nil {
				return 0, err
			}

			total += subSize
			continue
		}

		total += entryInfo.Size()
	}

	return total, nil
}

func FormatSize(size int64, isHuman bool) string {
	if !isHuman {
		return fmt.Sprintf("%dB", size)
	}

	const base int64 = 1024
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	if size < base {
		return fmt.Sprintf("%dB", size)
	}

	value := float64(size)
	unitIndex := 0
	for value >= float64(base) && unitIndex < len(units)-1 {
		value /= float64(base)
		unitIndex++
	}

	return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}
