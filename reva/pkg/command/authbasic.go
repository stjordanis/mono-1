package command

import (
	"context"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/cs3org/reva/cmd/revad/runtime"
	"github.com/gofrs/uuid"
	"github.com/micro/cli/v2"
	"github.com/oklog/run"
	"github.com/owncloud/mono/reva/pkg/config"
	"github.com/owncloud/mono/reva/pkg/flagset"
	"github.com/owncloud/mono/reva/pkg/server/debug"
)

// AuthBasic is the entrypoint for the auth-basic command.
func AuthBasic(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "auth-basic",
		Usage: "Start reva authprovider for basic auth",
		Flags: flagset.AuthBasicWithConfig(cfg),
		Before: func(c *cli.Context) error {
			cfg.Reva.AuthBasic.Services = c.StringSlice("service")

			return nil
		},
		Action: func(c *cli.Context) error {
			logger := NewLogger(cfg)

			if cfg.Tracing.Enabled {
				switch t := cfg.Tracing.Type; t {
				case "agent":
					logger.Error().
						Str("type", t).
						Msg("Reva only supports the jaeger tracing backend")

				case "jaeger":
					logger.Info().
						Str("type", t).
						Msg("configuring reva to use the jaeger tracing backend")

				case "zipkin":
					logger.Error().
						Str("type", t).
						Msg("Reva only supports the jaeger tracing backend")

				default:
					logger.Warn().
						Str("type", t).
						Msg("Unknown tracing backend")
				}

			} else {
				logger.Debug().
					Msg("Tracing is not enabled")
			}

			var (
				gr          = run.Group{}
				ctx, cancel = context.WithCancel(context.Background())
				//metrics     = metrics.New()
			)

			defer cancel()

			{

				uuid := uuid.Must(uuid.NewV4())
				pidFile := path.Join(os.TempDir(), "revad-"+c.Command.Name+"-"+uuid.String()+".pid")

				rcfg := map[string]interface{}{
					"core": map[string]interface{}{
						"max_cpus":             cfg.Reva.Users.MaxCPUs,
						"tracing_enabled":      cfg.Tracing.Enabled,
						"tracing_endpoint":     cfg.Tracing.Endpoint,
						"tracing_collector":    cfg.Tracing.Collector,
						"tracing_service_name": "auth-basic",
					},
					"shared": map[string]interface{}{
						"jwt_secret": cfg.Reva.JWTSecret,
					},
					"grpc": map[string]interface{}{
						"network": cfg.Reva.AuthBasic.Network,
						"address": cfg.Reva.AuthBasic.Addr,
						// TODO build services dynamically
						"services": map[string]interface{}{
							"authprovider": map[string]interface{}{
								"auth_manager": cfg.Reva.AuthProvider.Driver,
								"auth_managers": map[string]interface{}{
									"json": map[string]interface{}{
										"users": cfg.Reva.AuthProvider.JSON,
									},
									"ldap": map[string]interface{}{
										"hostname":      cfg.Reva.LDAP.Hostname,
										"port":          cfg.Reva.LDAP.Port,
										"base_dn":       cfg.Reva.LDAP.BaseDN,
										"loginfilter":   cfg.Reva.LDAP.LoginFilter,
										"bind_username": cfg.Reva.LDAP.BindDN,
										"bind_password": cfg.Reva.LDAP.BindPassword,
										"idp":           cfg.Reva.LDAP.IDP,
										"schema": map[string]interface{}{
											"dn":          "dn",
											"uid":         cfg.Reva.LDAP.Schema.UID,
											"mail":        cfg.Reva.LDAP.Schema.Mail,
											"displayName": cfg.Reva.LDAP.Schema.DisplayName,
											"cn":          cfg.Reva.LDAP.Schema.CN,
										},
									},
								},
							},
						},
					},
				}

				gr.Add(func() error {
					runtime.RunWithOptions(
						rcfg,
						pidFile,
						runtime.WithLogger(&logger.Logger),
					)
					return nil
				}, func(_ error) {
					logger.Info().
						Str("server", c.Command.Name).
						Msg("Shutting down server")

					cancel()
				})
			}

			{
				server, err := debug.Server(
					debug.Name(c.Command.Name+"-debug"),
					debug.Addr(cfg.Reva.AuthBasic.DebugAddr),
					debug.Logger(logger),
					debug.Context(ctx),
					debug.Config(cfg),
				)

				if err != nil {
					logger.Info().
						Err(err).
						Str("server", "debug").
						Msg("Failed to initialize server")

					return err
				}

				gr.Add(func() error {
					return server.ListenAndServe()
				}, func(_ error) {
					ctx, timeout := context.WithTimeout(ctx, 5*time.Second)

					defer timeout()
					defer cancel()

					if err := server.Shutdown(ctx); err != nil {
						logger.Info().
							Err(err).
							Str("server", "debug").
							Msg("Failed to shutdown server")
					} else {
						logger.Info().
							Str("server", "debug").
							Msg("Shutting down server")
					}
				})
			}

			{
				stop := make(chan os.Signal, 1)

				gr.Add(func() error {
					signal.Notify(stop, os.Interrupt)

					<-stop

					return nil
				}, func(err error) {
					close(stop)
					cancel()
				})
			}

			return gr.Run()
		},
	}
}
