package format

import "fmt"

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
