package init_

import "github.com/urfave/cli/v3"

var Command = &cli.Command{
	Name: "init",
	Flags: []cli.Flag{
		flagForce,
		flagType,
	},
	Usage:     "Initialize config file",
	ArgsUsage: `"global" | "local" | empty (defaults to "local")`,
	Action:    action,
}
