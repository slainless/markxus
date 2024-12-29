package init_

import "github.com/urfave/cli/v3"

var (
	flagForce = &cli.BoolFlag{
		Name:        "force",
		Aliases:     []string{"f"},
		Usage:       "Force config creation, overriding existing config without prompt, if exist",
		Destination: &initConfig.force,
		Value:       false,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("FORCE_OVERWRITE"),
		),
	}

	flagType = &cli.GenericFlag{
		Name:        "type",
		Aliases:     []string{"t"},
		Usage:       "Config type to be generated, either global or local. To be used with force flag",
		DefaultText: "global",
		Value:       initConfig.configType,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("FORCE_CONFIG_TYPE"),
		),
	}
)
