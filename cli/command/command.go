package command

import (
	"github.com/slainless/markxus/cli/config"
	"github.com/urfave/cli/v3"
)

var Main = &cli.Command{
	Flags: config.AllFlags,
}
