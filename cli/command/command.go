package command

import (
	"github.com/slainless/markxus/cli/markxus/command/config"
	"github.com/slainless/markxus/cli/markxus/command/generate"
	"github.com/urfave/cli/v3"
)

var Main = &cli.Command{
	Name:  "markxus",
	Usage: "Markxus CLI, an LLM-powered markdown converter for Nexus Mods page",
	Authors: []any{
		"slainless",
	},
	Commands: []*cli.Command{
		config.Command,
		generate.Command,
	},
}
