package config

import (
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
	"github.com/urfave/cli/v3"
)

var (
	LLMCategory = "LLM"

	FlagLlmProvider = &cli.GenericFlag{
		Name:        "llm-provider",
		Aliases:     []string{"gp"},
		Usage:       "LLM provider to be used, either from OpenAI (open_ai) or Google Generative AI (gen_ai)",
		DefaultText: "global",
		Category:    LLMCategory,
		Value:       &Config.Llm.Provider,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyLlmProvider),
			cli.NewMapValueSource(YamlKeyLlmProvider, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyLlmProvider, YamlSourceGlobal),
		),
	}

	FlagLlmApiKey = &cli.StringFlag{
		Name:        "llm-key",
		Aliases:     []string{"gk"},
		Destination: &Config.Llm.ApiKey,
		Category:    LLMCategory,
		Required:    true,
		Usage:       "API key to be used for AI generation",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyLlmApiKey),
			cli.NewMapValueSource(YamlKeyLlmApiKey, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyLlmApiKey, YamlSourceGlobal),
			cli.NewMapValueSource(YamlKeyLlmApiKey, KeyringSource),
		),
	}

	FlagLlmModelName = &cli.StringFlag{
		Name:        "model",
		Aliases:     []string{"m"},
		Destination: &Config.Llm.ModelName,
		Category:    LLMCategory,
		DefaultText: genai.DefaultModelName,
		Usage:       "Model name to be used for AI generation",
		Value:       genai.DefaultModelName,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyLlmModelName),
			cli.NewMapValueSource(YamlKeyLlmModelName, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyLlmModelName, YamlSourceGlobal),
		),
	}

	FlagMarkxusPromptFormat = &cli.StringFlag{
		Name:        "prompt",
		Aliases:     []string{"p"},
		Destination: &Config.Llm.Prompt,
		Category:    LLMCategory,
		DefaultText: "[[DefaultLlmPromptFormat]]",
		Usage:       "Prompt format to be used for AI generation",
		Value:       markxus.DefaultLlmPromptFormat,
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
			cli.NewMapValueSource(YamlKeyNexusApiKey, YamlSourceGlobal),
			cli.NewMapValueSource(YamlKeyNexusApiKey, KeyringSource),
		),
	}

	FlagNexusUrlGetModFormat = &cli.StringFlag{
		Name:        "api-mod-url-format",
		Aliases:     []string{"am"},
		Destination: &Config.Nexus.Url.GetModFormat,
		Category:    "Nexus",
		DefaultText: nexus.DefaultUrlGetModFormat,
		Usage:       "URL format for mod data API",
		Value:       nexus.DefaultUrlGetModFormat,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyNexusUrlGetModFormat),
			cli.NewMapValueSource(YamlKeyNexusUrlGetModFormat, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyNexusUrlGetModFormat, YamlSourceGlobal),
		),
	}

	FlagNexusUrlGetModFilesFormat = &cli.StringFlag{
		Name:        "api-files-url-format",
		Aliases:     []string{"af"},
		Destination: &Config.Nexus.Url.GetFilesFormat,
		Category:    "Nexus",
		DefaultText: nexus.DefaultUrlGetFilesFormat,
		Usage:       "URL format for mod files data API",
		Value:       nexus.DefaultUrlGetFilesFormat,
		Sources: cli.NewValueSourceChain(
			cli.EnvVar(EnvKeyNexusUrlGetFilesFormat),
			cli.NewMapValueSource(YamlKeyNexusUrlGetFilesFormat, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyNexusUrlGetFilesFormat, YamlSourceGlobal),
		),
	}

	FlagMarkxusUrlModPageFormat = &cli.StringFlag{
		Name:        "page-url-format",
		Aliases:     []string{"pf"},
		Destination: &Config.Nexus.Url.ModPageFormat,
		Category:    "Nexus",
		DefaultText: markxus.DefaultUrlModPageFormat,
		Usage:       "URL format to be for mod page",
		Value:       markxus.DefaultUrlModPageFormat,
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
		Value:       markxus.DefaultMarkdownHeaderFormat,
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
			cli.NewMapValueSource(YamlKeyOverwriteOutput, YamlSourceLocal),
			cli.NewMapValueSource(YamlKeyOverwriteOutput, YamlSourceGlobal),
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
	FlagLlmApiKey,
	FlagLlmModelName,
	FlagMarkxusPromptFormat,
	FlagNexusApiKey,
	FlagNexusUrlGetModFormat,
	FlagMarkxusUrlModPageFormat,
	FlagMarkdownHeaderFormat,
	FlagFallbackGameCode,
	FlagOverwriteOutput,
	FlagConfigType,
}
