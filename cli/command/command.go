package command

import (
	"github.com/slainless/markxus/cli/markxus/command/config"
	"github.com/urfave/cli/v3"
)

var Main = &cli.Command{
	Name:  "Markxus CLI",
	Usage: "LLM-powered markdown converter for Nexus Mods page",
	Commands: []*cli.Command{
		config.Command,
	},
}
