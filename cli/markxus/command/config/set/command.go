package set

import (
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:      "set",
	Aliases:   []string{"s"},
	Usage:     "Set configuration value",
	ArgsUsage: "<env_or_yaml_key> <value>",
	Flags: []cli.Flag{
		config.FlagConfigType,
	},
	Action: action,
}
