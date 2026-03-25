package code

import (
	"fmt"
	"os"
)

func GetSize(path string) (int64, error) {
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
		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		if entryInfo.IsDir() {
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
