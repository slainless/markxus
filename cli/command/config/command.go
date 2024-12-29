package config

import (
	"github.com/slainless/markxus/cli/command/config/cmd_init"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "config",
	Aliases: []string{"c", "cfg"},
	Commands: []*cli.Command{
		cmd_init.Command,
	},
	Usage: "Manage configuration",
}
