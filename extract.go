package testsuite

import (
	"context"

	"github.com/arduino/go-paths-helper"
	"github.com/codeclysm/extract/v3"
	"github.com/stretchr/testify/require"
)

// Extract extracts a tarball to a directory named as the archive
// with the "_content" suffix added. Returns the path to the directory.
func (e *Environment) Extract(archive *paths.Path) *paths.Path {
	destDir := archive.Parent().Join(archive.Base() + "_content")
	if destDir.Exist() {
		return destDir
	}

	t := e.T()

	file, err := archive.Open()
	require.NoError(t, err)
	defer file.Close()

	err = extract.Archive(context.Background(), file, destDir.String(), nil)
	require.NoError(t, err)

	return destDir
}
