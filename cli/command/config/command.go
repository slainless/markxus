package config

import (
	"github.com/slainless/markxus/cli/markxus/command/config/edit"
	"github.com/slainless/markxus/cli/markxus/command/config/init_"
	"github.com/slainless/markxus/cli/markxus/command/config/set"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "config",
	Aliases: []string{"c", "cfg"},
	Commands: []*cli.Command{
		init_.Command,
		set.Command,
		edit.Command,
	},
	Usage: "Manage configuration",
}
