package generate

import (
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/urfave/cli/v3"
)

var Command = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "Generate markdown from nexus mod",
	Action:  action,
	Flags: []cli.Flag{
		config.FlagFallbackGameCode,
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
