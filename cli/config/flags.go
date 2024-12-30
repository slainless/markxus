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
			cli.NewMapValueSource(YamlKeyGenAiApiKey, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyGenAiApiKey, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyGenAiModelName, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyGenAiModelName, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyMarkxusPromptFormat, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyMarkxusPromptFormat, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyNexusApiKey, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyGenAiApiKey, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyNexusUrlGetModFormat, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyNexusUrlGetModFormat, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyMarkxusUrlModPageFormat, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyMarkxusUrlModPageFormat, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyMarkdownHeaderFormat, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyMarkdownHeaderFormat, YamlSourceGlobal),
		),
	}

	FlagOutputDir = &cli.StringFlag{
		Name:        "output-dir",
		Aliases:     []string{"outdir", "o"},
		Destination: &Config.Generation.OutputDir,
		Category:    GenerationCategory,
		Usage:       "Output directory for markdown files",
		Value:       ".",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyOutputDir),
			cli.NewMapValueSource(YamlKeyOutputDir, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyOutputDir, YamlSourceGlobal),
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
			cli.NewMapValueSource(YamlKeyFallbackGameCode, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyFallbackGameCode, YamlSourceGlobal),
		),
	}
)

var (
	FlagOverwriteOutput = &cli.BoolFlag{
		Name:        "overwrite",
		Aliases:     []string{"w"},
		Usage:       "Overwrite existing markdown if exist",
		Value:       false,
		Destination: &Config.Common.Overwrite,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyOverwriteOutput),
			cli.NewMapValueSource(YamlKeyFallbackGameCode, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyFallbackGameCode, YamlSourceGlobal),
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
	FlagOverwriteOutput,
	FlagConfigType,
}
