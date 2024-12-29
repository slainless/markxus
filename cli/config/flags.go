package config

import (
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
	"github.com/urfave/cli/v3"
)

var (
	GenAICategory = "GenAI"

	GenAiApiKey = &cli.StringFlag{
		Name:        "genai-key",
		Aliases:     []string{"gk"},
		Destination: &Config.GenAi.ApiKey,
		Category:    GenAICategory,
		Required:    true,
		Usage:       "API key to be used for generative ai requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("GEN_AI_API_KEY"),
			cli.NewMapValueSource("genai_api_key", LocalYamlSource),
			cli.NewMapValueSource("genai_api_key", KeyringSource),
		),
	}

	GenAiModelName = &cli.StringFlag{
		Name:        "model",
		Aliases:     []string{"m"},
		Destination: &Config.GenAi.ModelName,
		Category:    GenAICategory,
		DefaultText: genai.DefaultModelName,
		Usage:       "Model name to be used for generative ai requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("GEN_AI_MODEL_NAME"),
			cli.NewMapValueSource("genai_model_name", LocalYamlSource),
			cli.NewMapValueSource("genai_model_name", GlobalYamlSource),
		),
	}

	MarkxusPromptFormat = &cli.StringFlag{
		Name:        "prompt",
		Aliases:     []string{"p"},
		Destination: &Config.GenAi.Prompt,
		Category:    GenAICategory,
		DefaultText: "[[DefaultGenAiPromptFormat]]",
		Usage:       "Prompt format to be used for generative ai requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("NEXUS_URL_GET_MOD_FORMAT"),
			cli.NewMapValueSource("nexus_url_get_mod_format", LocalYamlSource),
			cli.NewMapValueSource("nexus_url_get_mod_format", GlobalYamlSource),
		),
	}
)

var (
	NexusCategory = "Nexus"

	NexusApiKey = &cli.StringFlag{
		Name:        "nexus-key",
		Aliases:     []string{"nk"},
		Destination: &Config.Nexus.ApiKey,
		Category:    "Nexus",
		Required:    true,
		Usage:       "API key to be used for nexus requests",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("NEXUS_API_KEY"),
			cli.NewMapValueSource("nexus_api_key", LocalYamlSource),
			cli.NewMapValueSource("nexus_api_key", KeyringSource),
		),
	}

	NexusUrlGetModFormat = &cli.StringFlag{
		Name:        "api-url-format",
		Aliases:     []string{"af"},
		Destination: &Config.Nexus.Url.GetModFormat,
		Category:    "Nexus",
		DefaultText: nexus.DefaultUrlGetModFormat,
		Usage:       "URL format to be for mod data API",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("NEXUS_URL_GET_MOD_FORMAT"),
			cli.NewMapValueSource("nexus_url_get_mod_format", LocalYamlSource),
			cli.NewMapValueSource("nexus_url_get_mod_format", GlobalYamlSource),
		),
	}

	MarkxusUrlModPageFormat = &cli.StringFlag{
		Name:        "page-url-format",
		Aliases:     []string{"pf"},
		Destination: &Config.Nexus.Url.ModPageFormat,
		Category:    "Nexus",
		DefaultText: "[[DefaultUrlModPageFormat]]",
		Usage:       "URL format to be for mod page",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("NEXUS_URL_MOD_PAGE_FORMAT"),
			cli.NewMapValueSource("nexus_url_mod_page_format", LocalYamlSource),
			cli.NewMapValueSource("nexus_url_mod_page_format", GlobalYamlSource),
		),
	}
)

var (
	GenerationCategory = "Markdown"

	MarkdownHeaderFormat = &cli.StringFlag{
		Name:        "header-format",
		Aliases:     []string{"hf"},
		Destination: &Config.Generation.HeaderFormat,
		Category:    GenerationCategory,
		DefaultText: "[[DefaultMarkdownHeaderFormat]]",
		Usage:       "Template to be used for markdown header",

		Sources: cli.NewValueSourceChain(
			cli.EnvVar("MARKDOWN_HEADER_FORMAT"),
			cli.NewMapValueSource("markdown_header_format", LocalYamlSource),
			cli.NewMapValueSource("markdown_header_format", GlobalYamlSource),
		),
	}
)

var AllFlags = []cli.Flag{
	GenAiApiKey,
	GenAiModelName,
	MarkxusPromptFormat,
	NexusApiKey,
	NexusUrlGetModFormat,
	MarkxusUrlModPageFormat,
	MarkdownHeaderFormat,
}
