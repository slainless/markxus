package config

import (
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
	"github.com/urfave/cli/v3"
)

var (
	GenAICategory = "GenAI"

	FlagGenAiApiKey = &cli.StringFlag{
		Name:        "genai-key",
		Aliases:     []string{"gk"},
		Destination: &Config.GenAi.ApiKey,
		Category:    GenAICategory,
		Required:    true,
		Usage:       "API key to be used for generative ai requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyGenAiApiKey),
			cli.NewMapValueSource(YamlKeyGenAiApiKey, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyGenAiApiKey, KeyringSource),
		),
	}

	FlagGenAiModelName = &cli.StringFlag{
		Name:        "model",
		Aliases:     []string{"m"},
		Destination: &Config.GenAi.ModelName,
		Category:    GenAICategory,
		DefaultText: genai.DefaultModelName,
		Usage:       "Model name to be used for generative ai requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyGenAiModelName),
			cli.NewMapValueSource(YamlKeyGenAiModelName, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyGenAiModelName, GlobalYamlSource),
		),
	}

	FlagMarkxusPromptFormat = &cli.StringFlag{
		Name:        "prompt",
		Aliases:     []string{"p"},
		Destination: &Config.GenAi.Prompt,
		Category:    GenAICategory,
		DefaultText: "[[DefaultGenAiPromptFormat]]",
		Usage:       "Prompt format to be used for generative ai requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyMarkxusPromptFormat),
			cli.NewMapValueSource(YamlKeyMarkxusPromptFormat, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyMarkxusPromptFormat, GlobalYamlSource),
		),
	}
)

var (
	NexusCategory = "Nexus"

	FlagNexusApiKey = &cli.StringFlag{
		Name:        "nexus-key",
		Aliases:     []string{"nk"},
		Destination: &Config.Nexus.ApiKey,
		Category:    "Nexus",
		Required:    true,
		Usage:       "API key to be used for nexus requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyNexusApiKey),
			cli.NewMapValueSource(YamlKeyNexusApiKey, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyNexusApiKey, KeyringSource),
		),
	}

	FlagNexusUrlGetModFormat = &cli.StringFlag{
		Name:        "api-url-format",
		Aliases:     []string{"af"},
		Destination: &Config.Nexus.Url.GetModFormat,
		Category:    "Nexus",
		DefaultText: nexus.DefaultUrlGetModFormat,
		Usage:       "URL format to be for mod data API",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyNexusUrlGetModFormat),
			cli.NewMapValueSource(YamlKeyNexusUrlGetModFormat, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyNexusUrlGetModFormat, GlobalYamlSource),
		),
	}

	FlagMarkxusUrlModPageFormat = &cli.StringFlag{
		Name:        "page-url-format",
		Aliases:     []string{"pf"},
		Destination: &Config.Nexus.Url.ModPageFormat,
		Category:    "Nexus",
		DefaultText: markxus.DefaultUrlModPageFormat,
		Usage:       "URL format to be for mod page",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyMarkxusUrlModPageFormat),
			cli.NewMapValueSource(YamlKeyMarkxusUrlModPageFormat, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyMarkxusUrlModPageFormat, GlobalYamlSource),
		),
	}
)

var (
	GenerationCategory = "Markdown"

	FlagMarkdownHeaderFormat = &cli.StringFlag{
		Name:        "header-format",
		Aliases:     []string{"hf"},
		Destination: &Config.Generation.HeaderFormat,
		Category:    GenerationCategory,
		DefaultText: "[[DefaultMarkdownHeaderFormat]]",
		Usage:       "Template to be used for markdown header",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyMarkdownHeaderFormat),
			cli.NewMapValueSource(YamlKeyMarkdownHeaderFormat, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyMarkdownHeaderFormat, GlobalYamlSource),
		),
	}
)

var (
	HelperCategory = "Helper"

	FlagFallbackGameCode = &cli.StringFlag{
		Name:        "game-code",
		Aliases:     []string{"gc"},
		Destination: &Config.Helper.FallbackGameCode,
		Category:    HelperCategory,
		Usage:       "Fallback game code to use when no game code supplied in args",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyFallbackGameCode),
			cli.NewMapValueSource(YamlKeyFallbackGameCode, LocalYamlSource),
			cli.NewMapValueSource(YamlKeyFallbackGameCode, GlobalYamlSource),
		),
	}
)

var (
	FlagOverwrite = &cli.BoolFlag{
		Name:        "overwrite",
		Aliases:     []string{"w"},
		Usage:       "Overwrite file if exist",
		Value:       false,
		Destination: &Config.Common.Overwrite,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("OVERWRITE"),
		),
	}

	FlagConfigType = &cli.GenericFlag{
		Name:        "type",
		Aliases:     []string{"t"},
		Usage:       "Config type to be used, either global or local",
		DefaultText: "global",
		Value:       &Config.Common.ConfigType,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("CONFIG_TYPE"),
		),
	}
)

var AllFlags = []cli.Flag{
	FlagGenAiApiKey,
	FlagGenAiModelName,
	FlagMarkxusPromptFormat,
	FlagNexusApiKey,
	FlagNexusUrlGetModFormat,
	FlagMarkxusUrlModPageFormat,
	FlagMarkdownHeaderFormat,
	FlagFallbackGameCode,
	FlagOverwrite,
	FlagConfigType,
}
