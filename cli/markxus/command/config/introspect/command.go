package introspect

import (
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:        "introspect",
	Aliases:     []string{"spec"},
	Usage:       "Introspect configuration file",
	Description: "List resolved variables of the CLI from multiple sources, or from a single source (when `TYPE` is specified)",
	ArgsUsage:   `[TYPE ("global" | "local")]`,
	Action:      action,
	Flags: []cli.Flag{
		config.FlagFallbackGameCode,
		config.FlagGenAiProvider,
		config.FlagGenAiApiKey,
		config.FlagGenAiModelName,
		config.FlagMarkxusPromptFormat,
		config.FlagNexusApiKey,
		config.FlagNexusUrlGetModFormat,
		config.FlagMarkxusUrlModPageFormat,
		config.FlagMarkdownHeaderFormat,
		config.FlagOutputDir,
		config.FlagOverwriteOutput,
	},
}
