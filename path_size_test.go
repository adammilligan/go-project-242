package code

import (
	"fmt"
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
	type testCase struct {
		path      string
		recursive bool
		human     bool
		all       bool
		expected  string
		wantErr   bool
	}

	run := func(t *testing.T, name string, tc testCase) {
		t.Run(name, func(t *testing.T) {
			res, err := GetPathSize(tc.path, tc.recursive, tc.human, tc.all)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expected, res)
		})
	}

	t.Run("success scenarios", func(t *testing.T) {
		tests := map[string]testCase{
			"file: returns file size": {
				path:     fixturePath(t, "file.txt"),
				expected: fmt.Sprintf("%dB", fileSize(t, fixturePath(t, "file.txt"))),
			},
			"dir: first level size when recursive=false and all=false": {
				path: fixturePath(t, "dir"),
				expected: fmt.Sprintf(
					"%dB",
					fileSize(t, fixturePath(t, "dir", "a.txt"))+
						fileSize(t, fixturePath(t, "dir", "b.txt"))+
						fileSize(t, fixturePath(t, "dir", "hexlet.svg")),
				),
			},
		}

		for name, tc := range tests {
			run(t, name, tc)
		}
	})

	t.Run("error scenarios", func(t *testing.T) {
		tests := map[string]testCase{
			"missing path: returns error": {
				path:    fixturePath(t, "missing.txt"),
				wantErr: true,
			},
		}

		for name, tc := range tests {
			run(t, name, tc)
		}
	})
}
