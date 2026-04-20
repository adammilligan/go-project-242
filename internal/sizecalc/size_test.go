package sizecalc

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

	repoRoot := filepath.Join(filepath.Dir(thisFile), "..", "..")
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

func TestGetSize(t *testing.T) {
	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	svg := fixturePath(t, "dir", "hexlet.svg")
	hidden := fixturePath(t, "dir", ".hidden.txt")
	c := fixturePath(t, "dir", "sub", "c.txt")

	expectedFirstLevel := fileSize(t, a) + fileSize(t, b) + fileSize(t, svg)
	expectedRecursive := expectedFirstLevel + fileSize(t, c)
	expectedWithHidden := expectedFirstLevel + fileSize(t, hidden)

	tests := map[string]struct {
		path      string
		all       bool
		recursive bool
		expected  int64
		wantErr   bool
	}{
		"file: returns file size": {
			path:     fixturePath(t, "file.txt"),
			expected: fileSize(t, fixturePath(t, "file.txt")),
		},
		"file: hidden file returns error when all=false": {
			path:    fixturePath(t, "dir", ".hidden.txt"),
			all:     false,
			wantErr: true,
		},
		"file: hidden file returns size when all=true": {
			path:     fixturePath(t, "dir", ".hidden.txt"),
			all:      true,
			expected: fileSize(t, fixturePath(t, "dir", ".hidden.txt")),
		},
		"dir: first level only when recursive=false": {
			path:     fixturePath(t, "dir"),
			all:      false,
			expected: expectedFirstLevel,
		},
		"dir: includes nested directories when recursive=true": {
			path:      fixturePath(t, "dir"),
			all:       false,
			recursive: true,
			expected:  expectedRecursive,
		},
		"dir: hidden entries are excluded when all=false": {
			path:     fixturePath(t, "dir"),
			all:      false,
			expected: expectedFirstLevel,
		},
		"dir: hidden entries are included when all=true": {
			path:     fixturePath(t, "dir"),
			all:      true,
			expected: expectedWithHidden,
		},
		"missing path: returns error": {
			path:    fixturePath(t, "missing.txt"),
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetSize(tc.path, tc.all, tc.recursive)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
