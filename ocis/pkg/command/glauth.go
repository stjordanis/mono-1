package command

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/mono/glauth/pkg/command"
	svcconfig "github.com/owncloud/mono/glauth/pkg/config"
	"github.com/owncloud/mono/glauth/pkg/flagset"
	"github.com/owncloud/mono/ocis/pkg/config"
	"github.com/owncloud/mono/ocis/pkg/register"
)

// GLAuthCommand is the entrypoint for the glauth command.
func GLAuthCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "glauth",
		Usage:    "Start glauth server",
		Category: "Extensions",
		Flags:    flagset.ServerWithConfig(cfg.GLAuth),
		Action: func(c *cli.Context) error {
			scfg := configureGLAuth(cfg)

			return cli.HandleAction(
				command.Server(scfg).Action,
				c,
			)
		},
	}
}

func configureGLAuth(cfg *config.Config) *svcconfig.Config {
	cfg.GLAuth.Log.Level = cfg.Log.Level
	cfg.GLAuth.Log.Pretty = cfg.Log.Pretty
	cfg.GLAuth.Log.Color = cfg.Log.Color

	if cfg.Tracing.Enabled {
		cfg.GLAuth.Tracing.Enabled = cfg.Tracing.Enabled
		cfg.GLAuth.Tracing.Type = cfg.Tracing.Type
		cfg.GLAuth.Tracing.Endpoint = cfg.Tracing.Endpoint
		cfg.GLAuth.Tracing.Collector = cfg.Tracing.Collector
		cfg.GLAuth.Tracing.Service = cfg.Tracing.Service
	}

	return cfg.GLAuth
}

func init() {
	register.AddCommand(GLAuthCommand)
}
