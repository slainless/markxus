package config

const Version = "v0.1.9"

type MarkxusConfig struct {
	Llm struct {
		Provider  EnumValue[Provider]
		ApiKey    string
		ModelName string

		Prompt string
	}

	Nexus struct {
		ApiKey string

		Url struct {
			GetModFormat  string
			ModPageFormat string
		}
	}

	Generation struct {
		HeaderFormat string
		OutputDir    string
	}

	Helper struct {
		FallbackGameCode string
		Interactive      bool
	}

	Common struct {
		ConfigType EnumValue[ConfigType]
		Overwrite  bool
	}
}

var Config MarkxusConfig

func init() {
	Config.Llm.Provider = EnumValue[Provider]{
		Enum:    []Provider{ProviderOpenAi, ProviderGenAi},
		Default: ProviderGenAi,
	}
	Config.Common.ConfigType = EnumValue[ConfigType]{
		Enum:    []ConfigType{ConfigTypeGlobal, ConfigTypeLocal},
		Default: ConfigTypeGlobal,
	}
}

var (
	YamlKeyLlmProvider             = "llm_provider"
	YamlKeyLlmApiKey               = "llm_api_key"
	YamlKeyLlmModelName            = "llm_model_name"
	YamlKeyMarkxusPromptFormat     = "llm_prompt_format"
	YamlKeyNexusApiKey             = "nexus_api_key"
	YamlKeyNexusUrlGetModFormat    = "nexus_url_get_mod_format"
	YamlKeyMarkxusUrlModPageFormat = "nexus_url_mod_page_format"
	YamlKeyMarkdownHeaderFormat    = "markdown_header_format"
	YamlKeyFallbackGameCode        = "fallback_game_code"
	YamlKeyOutputDir               = "output_dir"
	YamlKeyOverwriteOutput         = "overwrite_output"
	YamlKeyCategoryIconMap         = "category_icon_map"
)

var (
	EnvKeyLlmProvider             = "LLM_PROVIDER"
	EnvKeyLlmApiKey               = "LLM_API_KEY"
	EnvKeyLlmModelName            = "LLM_MODEL_NAME"
	EnvKeyMarkxusPromptFormat     = "LLM_PROMPT_FORMAT"
	EnvKeyNexusApiKey             = "NEXUS_API_KEY"
	EnvKeyNexusUrlGetModFormat    = "NEXUS_URL_GET_MOD_FORMAT"
	EnvKeyMarkxusUrlModPageFormat = "NEXUS_URL_MOD_PAGE_FORMAT"
	EnvKeyMarkdownHeaderFormat    = "MARKDOWN_HEADER_FORMAT"
	EnvKeyFallbackGameCode        = "FALLBACK_GAME_CODE"
	EnvKeyOutputDir               = "OUTPUT_DIR"
	EnvKeyOverwriteOutput         = "OVERWRITE_OUTPUT"
)

var yamlToEnv = map[string]string{
	YamlKeyLlmProvider:             EnvKeyLlmProvider,
	YamlKeyLlmApiKey:               EnvKeyLlmApiKey,
	YamlKeyLlmModelName:            EnvKeyLlmModelName,
	YamlKeyMarkxusPromptFormat:     EnvKeyMarkxusPromptFormat,
	YamlKeyNexusApiKey:             EnvKeyNexusApiKey,
	YamlKeyNexusUrlGetModFormat:    EnvKeyNexusUrlGetModFormat,
	YamlKeyMarkxusUrlModPageFormat: EnvKeyMarkxusUrlModPageFormat,
	YamlKeyMarkdownHeaderFormat:    EnvKeyMarkdownHeaderFormat,
	YamlKeyFallbackGameCode:        EnvKeyFallbackGameCode,
	YamlKeyOutputDir:               EnvKeyOutputDir,
	YamlKeyOverwriteOutput:         EnvKeyOverwriteOutput,
}

var envToYaml = map[string]string{
	EnvKeyLlmProvider:             YamlKeyLlmProvider,
	EnvKeyLlmApiKey:               YamlKeyLlmApiKey,
	EnvKeyLlmModelName:            YamlKeyLlmModelName,
	EnvKeyMarkxusPromptFormat:     YamlKeyMarkxusPromptFormat,
	EnvKeyNexusApiKey:             YamlKeyNexusApiKey,
	EnvKeyNexusUrlGetModFormat:    YamlKeyNexusUrlGetModFormat,
	EnvKeyMarkxusUrlModPageFormat: YamlKeyMarkxusUrlModPageFormat,
	EnvKeyMarkdownHeaderFormat:    YamlKeyMarkdownHeaderFormat,
	EnvKeyFallbackGameCode:        YamlKeyFallbackGameCode,
	EnvKeyOutputDir:               YamlKeyOutputDir,
	EnvKeyOverwriteOutput:         YamlKeyOverwriteOutput,
}

func EnvToYaml(key string) string {
	return envToYaml[key]
}

func YamlToEnv(key string) string {
	return yamlToEnv[key]
}

func Resolve(envKey string) any {
	switch envKey {
	case EnvKeyLlmProvider:
		return Config.Llm.Provider.Selected()
	case EnvKeyLlmApiKey:
		return Config.Llm.ApiKey
	case EnvKeyLlmModelName:
		return Config.Llm.ModelName
	case EnvKeyMarkxusPromptFormat:
		return Config.Llm.Prompt
	case EnvKeyNexusApiKey:
		return Config.Nexus.ApiKey
	case EnvKeyNexusUrlGetModFormat:
		return Config.Nexus.Url.GetModFormat
	case EnvKeyMarkxusUrlModPageFormat:
		return Config.Nexus.Url.ModPageFormat
	case EnvKeyMarkdownHeaderFormat:
		return Config.Generation.HeaderFormat
	case EnvKeyFallbackGameCode:
		return Config.Helper.FallbackGameCode
	case EnvKeyOutputDir:
		return Config.Generation.OutputDir
	case EnvKeyOverwriteOutput:
		return Config.Common.Overwrite
	}

	return nil
}

type ConfigType string

var (
	ConfigTypeGlobal ConfigType = "global"
	ConfigTypeLocal  ConfigType = "local"
)

type Provider string

var (
	ProviderOpenAi Provider = "open_ai"
	ProviderGenAi  Provider = "gen_ai"
)
