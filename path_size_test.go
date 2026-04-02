package code

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func fixturePath(t *testing.T, parts ...string) string {
	t.Helper()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve runtime caller")
	}

	repoRoot := filepath.Dir(thisFile)
	allParts := append([]string{repoRoot, "testdata"}, parts...)
	return filepath.Join(allParts...)
}

func fileSize(t *testing.T, path string) int64 {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}

	return info.Size()
}

func TestGetPathSize(t *testing.T) {
	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	svg := fixturePath(t, "dir", "hexlet.svg")
	hidden := fixturePath(t, "dir", ".hidden.txt")
	c := fixturePath(t, "dir", "sub", "c.txt")

	expectedFirstLevel := fileSize(t, a) + fileSize(t, b) + fileSize(t, svg)
	expectedRecursive := expectedFirstLevel + fileSize(t, c)
	expectedWithHidden := expectedFirstLevel + fileSize(t, hidden)

	tests := []struct {
		name        string
		path        string
		recursive   bool
		human       bool
		all         bool
		expected    string
		expectedNot string
		wantErr     bool
	}{
		{
			name:     "file",
			path:     fixturePath(t, "file.txt"),
			expected: FormatSize(fileSize(t, fixturePath(t, "file.txt")), false),
		},
		{
			name:        "directory-first-level",
			path:        fixturePath(t, "dir"),
			recursive:   false,
			human:       false,
			all:         false,
			expected:    FormatSize(expectedFirstLevel, false),
			expectedNot: FormatSize(expectedRecursive, false),
		},
		{
			name:        "directory-first-level-hidden-all-false",
			path:        fixturePath(t, "dir"),
			all:         false,
			expected:    FormatSize(expectedFirstLevel, false),
			expectedNot: FormatSize(expectedWithHidden, false),
		},
		{
			name:     "directory-first-level-hidden-all-true",
			path:     fixturePath(t, "dir"),
			all:      true,
			expected: FormatSize(expectedWithHidden, false),
		},
		{
			name:      "directory-recursive",
			path:      fixturePath(t, "dir"),
			recursive: true,
			all:       false,
			expected:  FormatSize(expectedRecursive, false),
		},
		{
			name:    "missing-path",
			path:    fixturePath(t, "missing.txt"),
			wantErr: true,
		},
		{
			name:     "directory-first-level-human",
			path:     fixturePath(t, "dir"),
			all:      false,
			human:    true,
			expected: FormatSize(expectedFirstLevel, true),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := GetPathSize(tc.path, tc.recursive, tc.human, tc.all)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expected, res)
			if tc.expectedNot != "" {
				require.NotEqual(t, tc.expectedNot, res)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		{
			name:     "bytes",
			size:     123,
			human:    false,
			expected: "123B",
		},
		{
			name:     "bytes-big",
			size:     25165824,
			human:    false,
			expected: "25165824B",
		},
		{
			name:     "human-mb",
			size:     25165824,
			human:    true,
			expected: "24.0MB",
		},
		{
			name:     "human-mb-smaller",
			size:     1234567,
			human:    true,
			expected: "1.2MB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, FormatSize(tc.size, tc.human))
		})
	}
}
