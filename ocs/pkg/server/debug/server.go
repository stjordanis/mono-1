package debug

import (
	"io"
	"net/http"

	"github.com/owncloud/mono/ocs/pkg/config"
	"github.com/owncloud/mono/ocs/pkg/version"
	"github.com/owncloud/mono/ocis-pkg/service/debug"
)

// Server initializes the debug service and server.
func Server(opts ...Option) (*http.Server, error) {
	options := newOptions(opts...)

	return debug.NewService(
		debug.Logger(options.Logger),
		debug.Name("ocs"),
		debug.Version(version.String),
		debug.Address(options.Config.Debug.Addr),
		debug.Token(options.Config.Debug.Token),
		debug.Pprof(options.Config.Debug.Pprof),
		debug.Zpages(options.Config.Debug.Zpages),
		debug.Health(health(options.Config)),
		debug.Ready(ready(options.Config)),
	), nil
}

// health implements the health check.
func health(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		// TODO(tboerger): check if services are up and running

		io.WriteString(w, http.StatusText(http.StatusOK))
	}
}

// ready implements the ready check.
func ready(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		// TODO(tboerger): check if services are up and running

		io.WriteString(w, http.StatusText(http.StatusOK))
	}
}
