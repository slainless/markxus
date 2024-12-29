package cmd_init

import "github.com/urfave/cli/v3"

var Command = &cli.Command{
	Name: "init",
	Flags: []cli.Flag{
		flagForce,
	},
	Usage:     "Initialize config file",
	ArgsUsage: `"global" | "local" | empty (defaults to "local")`,
	Action:    action,
}
