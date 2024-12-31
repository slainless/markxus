package init_

import (
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Flags: []cli.Flag{
		config.FlagConfigType,
		&cli.BoolFlag{
			Name:        "overwrite",
			Aliases:     []string{"w"},
			Usage:       "Overwrite existing config file if exist",
			Value:       false,
			Destination: &config.Config.Common.Overwrite,
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("OVERWRITE_CONFIG"),
			),
		},
	},
	Usage:  "Initialize config file",
	Action: action,
}
