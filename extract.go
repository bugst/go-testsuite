package testsuite

import (
	"context"

	"github.com/arduino/go-paths-helper"
	"github.com/codeclysm/extract/v3"
)

// Extract extracts a tarball to a directory named as the archive
// with the "_content" suffix added. Returns the path to the directory.
func (e *Environment) Extract(archive *paths.Path) *paths.Path {
	destDir := archive.Parent().Join(archive.Base() + "_content")
	if destDir.Exist() {
		return destDir
	}

	file, err := archive.Open()
	e.t.NoError(err)
	defer file.Close()

	err = extract.Archive(context.Background(), file, destDir.String(), nil)
	e.t.NoError(err)

	return destDir
}
