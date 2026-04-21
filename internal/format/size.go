package format

import (
	"errors"
	"fmt"
)

var ErrNegativeSize = errors.New("size must be non-negative")

// FormatSize formats a size in bytes.
//
// When isHuman is false, it returns bytes (e.g. "123B").
// When isHuman is true, it uses base-1024 units (e.g. "1.2MB").
func FormatSize(size int64, isHuman bool) (string, error) {
	if size < 0 {
		return "", fmt.Errorf("%w: got %d", ErrNegativeSize, size)
	}

	const base int64 = 1024

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	if !isHuman || size < base {
		return fmt.Sprintf("%d%s", size, units[0]), nil
	}

	value := float64(size)
	baseFloat := float64(base)

	unit := units[0]

	for _, u := range units[1:] {
		if value < baseFloat {
			break
		}

		value /= baseFloat
		unit = u
	}

	return fmt.Sprintf("%.1f%s", value, unit), nil
}
