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

	res, err := code.GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Equal(t, code.FormatSize(expected, false), res)
}

func TestGetPathSize_DirectoryFirstLevel(t *testing.T) {
	dir := fixturePath(t, "dir")

	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	svg := fixturePath(t, "dir", "hexlet.svg")
	c := fixturePath(t, "dir", "sub", "c.txt")

	expected := fileSize(t, a) + fileSize(t, b) + fileSize(t, svg)
	expectedNested := expected + fileSize(t, c)

	res, err := code.GetPathSize(dir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, code.FormatSize(expected, false), res)
	require.NotEqual(t, code.FormatSize(expectedNested, false), res)
}

func TestGetPathSize_DirectoryFirstLevel_HiddenFiles(t *testing.T) {
	dir := fixturePath(t, "dir")

	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	svg := fixturePath(t, "dir", "hexlet.svg")
	hidden := fixturePath(t, "dir", ".hidden.txt")

	expected := fileSize(t, a) + fileSize(t, b) + fileSize(t, svg)
	expectedWithHidden := expected + fileSize(t, hidden)

	resWithoutAll, err := code.GetPathSize(dir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, code.FormatSize(expected, false), resWithoutAll)

	resWithAll, err := code.GetPathSize(dir, false, false, true)
	require.NoError(t, err)
	require.Equal(t, code.FormatSize(expectedWithHidden, false), resWithAll)

	// human=false: сравнение по байтам строково эквивалентно числам.
	require.NotEqual(t, resWithAll, "")
	require.GreaterOrEqual(t, expectedWithHidden, expected)
}

func TestGetPathSize_DirectoryRecursive(t *testing.T) {
	dir := fixturePath(t, "dir")

	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	svg := fixturePath(t, "dir", "hexlet.svg")
	c := fixturePath(t, "dir", "sub", "c.txt")

	expected := fileSize(t, a) + fileSize(t, b) + fileSize(t, svg) + fileSize(t, c)

	res, err := code.GetPathSize(dir, true, false, false)
	require.NoError(t, err)
	require.Equal(t, code.FormatSize(expected, false), res)
}

func TestGetPathSize_DirectoryFirstLevel_HumanReadable(t *testing.T) {
	dir := fixturePath(t, "dir")

	a := fixturePath(t, "dir", "a.txt")
	b := fixturePath(t, "dir", "b.txt")
	svg := fixturePath(t, "dir", "hexlet.svg")

	expected := fileSize(t, a) + fileSize(t, b) + fileSize(t, svg)

	res, err := code.GetPathSize(dir, false, true, false)
	require.NoError(t, err)
	require.Equal(t, code.FormatSize(expected, true), res)
}

func TestGetPathSize_MissingPath(t *testing.T) {
	path := fixturePath(t, "missing.txt")

	_, err := code.GetPathSize(path, false, false, false)
	require.Error(t, err)
}

func TestFormatSize(t *testing.T) {
	require.Equal(t, "123B", code.FormatSize(123, false))
	require.Equal(t, "25165824B", code.FormatSize(25165824, false))

	require.Equal(t, "24.0MB", code.FormatSize(25165824, true))
	require.Equal(t, "1.2MB", code.FormatSize(1234567, true))
}
