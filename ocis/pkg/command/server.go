// +build !simple

package command

import (
	"strings"

	"github.com/micro/cli/v2"
	"github.com/owncloud/mono/ocis/pkg/config"
	"github.com/owncloud/mono/ocis/pkg/flagset"
	"github.com/owncloud/mono/ocis/pkg/register"
	"github.com/owncloud/mono/ocis/pkg/runtime"
	"github.com/owncloud/mono/ocis/pkg/tracing"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    "Start fullstack server",
		Category: "Fullstack",
		Flags:    flagset.ServerWithConfig(cfg),
		Before: func(c *cli.Context) error {
			if cfg.HTTP.Root != "/" {
				cfg.HTTP.Root = strings.TrimSuffix(cfg.HTTP.Root, "/")
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			if err := tracing.Start(cfg); err != nil {
				return err
			}

			r := runtime.New()
			// TODO temporary service startup selection. Should go away and the runtime should take care of it.
			return r.Start(append([]string{
				"accounts",
				"settings",
				"konnectd",
				"proxy",
				"phoenix",
				"reva-frontend",
				"reva-gateway",
			}, runtime.MicroServices...)...)
		},
	}
}

func init() {
	register.AddCommand(Server)
}
