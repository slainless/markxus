package init_

import (
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Flags: []cli.Flag{
		config.FlagOverwrite,
		config.FlagConfigType,
	},
	Usage:  "Initialize config file",
	Action: action,
}
