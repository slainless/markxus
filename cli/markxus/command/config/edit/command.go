package edit

import (
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "edit",
	Aliases: []string{"e"},
	Usage:   "Open configuration file with default editor",
	Action:  action,
	Flags: []cli.Flag{
		config.FlagConfigType,
		&cli.BoolFlag{
			Name:        "init",
			Aliases:     []string{"i"},
			Usage:       "Initialize config file if not exist",
			Value:       false,
			Destination: &config.Config.Common.Overwrite,
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("INIT_CONFIG"),
			),
		},
	},
}
