package debug

import (
	"io"
	"net/http"

	"github.com/owncloud/mono/ocis-pkg/service/debug"
	"github.com/owncloud/mono/thumbnails/pkg/config"
	"github.com/owncloud/mono/thumbnails/pkg/version"
)

// Server initializes the debug service and server.
func Server(opts ...Option) (*http.Server, error) {
	options := newOptions(opts...)

	return debug.NewService(
		debug.Logger(options.Logger),
		debug.Name(options.Config.Server.Name),
		debug.Version(version.String),
		debug.Address(options.Config.Debug.Addr),
		debug.Token(options.Config.Debug.Token),
		debug.Pprof(options.Config.Debug.Pprof),
		debug.Zpages(options.Config.Debug.Zpages),
		debug.Health(health(options.Config)),
		debug.Ready(ready(options.Config)),
	), nil
}

func health(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, http.StatusText(http.StatusOK))
	}
}

func ready(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, http.StatusText(http.StatusOK))
	}
}
