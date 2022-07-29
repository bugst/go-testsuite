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
	}
}

// CleanUp removes the test environment.
func (e *Environment) CleanUp() {
	e.t.NoError(e.rootDir.RemoveAll())
}

// Root returns the root dir of the environment.
func (e *Environment) Root() *paths.Path {
	return e.rootDir
}
