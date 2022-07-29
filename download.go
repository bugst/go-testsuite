package testsuite

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/arduino/go-paths-helper"
	"github.com/stretchr/testify/require"
)

// Download downloads a file from a URL and returns the path to the downloaded file.
// The file is saved and cached in a shared downloads directory.
// If the file already exists, it is not downloaded again.
func (e *Environment) Download(rawURL string) *paths.Path {
	t := e.T()

	url, err := url.Parse(rawURL)
	require.NoError(t, err)

	filename := filepath.Base(url.Path)
	if filename == "/" {
		filename = ""
	} else {
		filename = "-" + filename
	}

	hash := md5.Sum([]byte(rawURL))
	resource := e.downloadsDir.Join(hex.EncodeToString(hash[:]) + filename)

	// If the resource already exist, return it
	if resource.Exist() {
		return resource
	}

	// Download file
	resp, err := http.Get(rawURL)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Copy data in a temp file
	tmp := resource.Parent().Join(resource.Base() + ".tmp")
	out, err := tmp.Create()
	require.NoError(t, err)
	defer tmp.Remove()
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	require.NoError(t, err)
	require.NoError(t, out.Close())

	// Rename the file to its final destination
	require.NoError(t, tmp.Rename(resource))

	return resource
}
