package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize возвращает размер указанного пути в виде отформатированной строки.
//
// Если path указывает на файл, возвращается размер файла.
// Если path указывает на директорию, размер директории рассчитывается по её содержимому.
// Если recursive=true, учитываются размеры вложенных директорий.
// Если all=false, элементы с именами, начинающимися с '.', пропускаются.
// Если human=true, размер выводится в человекочитаемых двоичных единицах (основание 1024).
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, all, recursive)
	if err != nil {
		return "", fmt.Errorf("get path size for %q: %w", path, err)
	}

	return FormatSize(size, human), nil
}

// GetSize вычисляет общий размер в байтах для указанного пути.
//
// Если path указывает на файл, возвращается размер файла.
// Если path указывает на директорию, суммируются размеры файлов внутри неё.
// Если recursive=true, учитываются размеры файлов во вложенных директориях.
// Если all=false, элементы с именами, начинающимися с '.', пропускаются.
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

// FormatSize преобразует размер в байтах в строковое представление.
//
// Если isHuman=false, выводится размер в байтах (например, "123B").
// Если isHuman=true, выводится размер в человекочитаемых двоичных единицах (основание 1024),
// например "1.2MB".
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
