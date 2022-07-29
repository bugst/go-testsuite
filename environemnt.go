package testsuite

import (
	"testing"

	"github.com/arduino/go-paths-helper"
	"github.com/stretchr/testify/require"
)

// ProjectName is the prefix used in the test temp files
var ProjectName = "go-testsuite"

// Environment is a test environment for the test suite.
type Environment struct {
	rootDir      *paths.Path
	downloadsDir *paths.Path
	t            *require.Assertions
	cleanUp      func()
}

// SharedDir returns the shared downloads directory.
func SharedDir(t *testing.T, id string) *paths.Path {
	downloadsDir := paths.TempDir().Join(ProjectName + "-" + id)
	require.NoError(t, downloadsDir.MkdirAll())
	return downloadsDir
}

// NewEnvironment creates a new test environment.
func NewEnvironment(t *testing.T) *Environment {
	downloadsDir := SharedDir(t, "downloads")
	rootDir, err := paths.MkTempDir("", ProjectName)
	require.NoError(t, err)
	return &Environment{
		rootDir:      rootDir,
		downloadsDir: downloadsDir,
		t:            require.New(t),
		cleanUp: func() {
			require.NoError(t, rootDir.RemoveAll())
		},
	}
}

// RegisterCleanUpCallback adds a clean up function to the clean up chain
func (e *Environment) RegisterCleanUpCallback(newCleanUp func()) {
	previousCleanUp := e.cleanUp
	e.cleanUp = func() {
		newCleanUp()
		previousCleanUp()
	}
}

// CleanUp removes the test environment.
func (e *Environment) CleanUp() {
	e.cleanUp()
}

// RootDir returns the root dir of the environment.
func (e *Environment) RootDir() *paths.Path {
	return e.rootDir
}

// SharedDownloadsDir return the shared directory for downloads
func (e *Environment) SharedDownloadsDir() *paths.Path {
	return e.downloadsDir
}
