package config

import (
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
			cli.EnvVar("GEN_AI_API_KEY"),
			cli.NewMapValueSource(KeyGenAiApiKey, LocalYamlSource),
			cli.NewMapValueSource(KeyGenAiApiKey, KeyringSource),
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
			cli.EnvVar("GEN_AI_MODEL_NAME"),
			cli.NewMapValueSource(KeyGenAiModelName, LocalYamlSource),
			cli.NewMapValueSource(KeyGenAiModelName, GlobalYamlSource),
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
			cli.EnvVar("NEXUS_URL_GET_MOD_FORMAT"),
			cli.NewMapValueSource(KeyMarkxusPromptFormat, LocalYamlSource),
			cli.NewMapValueSource(KeyMarkxusPromptFormat, GlobalYamlSource),
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
			cli.EnvVar("NEXUS_API_KEY"),
			cli.NewMapValueSource(KeyNexusApiKey, LocalYamlSource),
			cli.NewMapValueSource(KeyNexusApiKey, KeyringSource),
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
			cli.EnvVar("NEXUS_URL_GET_MOD_FORMAT"),
			cli.NewMapValueSource(KeyNexusUrlGetModFormat, LocalYamlSource),
			cli.NewMapValueSource(KeyNexusUrlGetModFormat, GlobalYamlSource),
		),
	}

	FlagMarkxusUrlModPageFormat = &cli.StringFlag{
		Name:        "page-url-format",
		Aliases:     []string{"pf"},
		Destination: &Config.Nexus.Url.ModPageFormat,
		Category:    "Nexus",
		DefaultText: "[[DefaultUrlModPageFormat]]",
		Usage:       "URL format to be for mod page",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("NEXUS_URL_MOD_PAGE_FORMAT"),
			cli.NewMapValueSource(KeyMarkxusUrlModPageFormat, LocalYamlSource),
			cli.NewMapValueSource(KeyMarkxusUrlModPageFormat, GlobalYamlSource),
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
			cli.EnvVar("MARKDOWN_HEADER_FORMAT"),
			cli.NewMapValueSource(KeyMarkdownHeaderFormat, LocalYamlSource),
			cli.NewMapValueSource(KeyMarkdownHeaderFormat, GlobalYamlSource),
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
}
