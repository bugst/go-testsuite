package testsuite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/arduino/go-paths-helper"
	"github.com/stretchr/testify/require"
)

// HTTPServeFile spawn an http server that serve a single file. The server
// is started on the given port. The URL to the file and a cleanup function are returned.
func (env *Environment) HTTPServeFile(port uint16, path *paths.Path) *url.URL {
	mux := http.NewServeMux()
	mux.HandleFunc("/"+path.Base(), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.String())
	})
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	t := env.T()
	fileURL, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%d/%s", port, path.Base()))
	require.NoError(t, err)

	go func() {
		err := server.ListenAndServe()
		require.Equal(t, err, http.ErrServerClosed)
	}()

	env.RegisterCleanUpCallback(func() {
		server.Close()
	})

	return fileURL
}
