package code

import "os"

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
