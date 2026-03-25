package tests

import (
	"code"
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

	repoRoot := filepath.Join(filepath.Dir(thisFile), "..")
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

func TestGetPathSize_File(t *testing.T) {
	path := fixturePath(t, "file.txt")
	expected := fileSize(t, path)

	size, err := code.GetSize(path)
	require.NoError(t, err)
	require.Equal(t, expected, size)
}

func TestGetPathSize_DirectoryFirstLevel(t *testing.T) {
	dir := fixturePath(t, "dir")

	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	c := fixturePath(t, "dir", "sub", "c.txt")

	expected := fileSize(t, a) + fileSize(t, b)
	expectedNested := expected + fileSize(t, c)

	size, err := code.GetSize(dir)
	require.NoError(t, err)
	require.Equal(t, expected, size)
	require.NotEqual(t, expectedNested, size)
}

func TestGetPathSize_MissingPath(t *testing.T) {
	path := fixturePath(t, "missing.txt")

	_, err := code.GetSize(path)
	require.Error(t, err)
}

func TestFormatSize(t *testing.T) {
	require.Equal(t, "123B", code.FormatSize(123, false))
	require.Equal(t, "25165824B", code.FormatSize(25165824, false))

	require.Equal(t, "24.0MB", code.FormatSize(25165824, true))
	require.Equal(t, "1.2MB", code.FormatSize(1234567, true))
}
