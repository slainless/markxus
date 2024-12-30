package config

type MarkxusConfig struct {
	GenAi struct {
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
	Config.Common.ConfigType = EnumValue[ConfigType]{
		Enum:    []ConfigType{ConfigTypeGlobal, ConfigTypeLocal},
		Default: ConfigTypeGlobal,
	}
}

var (
	YamlKeyGenAiApiKey             = "genai_api_key"
	YamlKeyGenAiModelName          = "genai_model_name"
	YamlKeyMarkxusPromptFormat     = "genai_prompt_format"
	YamlKeyNexusApiKey             = "nexus_api_key"
	YamlKeyNexusUrlGetModFormat    = "nexus_url_get_mod_format"
	YamlKeyMarkxusUrlModPageFormat = "nexus_url_mod_page_format"
	YamlKeyMarkdownHeaderFormat    = "markdown_header_format"
	YamlKeyFallbackGameCode        = "fallback_game_code"
	YamlKeyInteractive             = "interactive_cli"
)

var (
	EnvKeyGenAiApiKey             = "GEN_AI_API_KEY"
	EnvKeyGenAiModelName          = "GEN_AI_MODEL_NAME"
	EnvKeyMarkxusPromptFormat     = "GEN_AI_PROMPT_FORMAT"
	EnvKeyNexusApiKey             = "NEXUS_API_KEY"
	EnvKeyNexusUrlGetModFormat    = "NEXUS_URL_GET_MOD_FORMAT"
	EnvKeyMarkxusUrlModPageFormat = "NEXUS_URL_MOD_PAGE_FORMAT"
	EnvKeyMarkdownHeaderFormat    = "MARKDOWN_HEADER_FORMAT"
	EnvKeyFallbackGameCode        = "FALLBACK_GAME_CODE"
	EnvKeyInteractive             = "INTERACTIVE_CLI"
)

var yamlToEnv = map[string]string{
	YamlKeyGenAiApiKey:             EnvKeyGenAiApiKey,
	YamlKeyGenAiModelName:          EnvKeyGenAiModelName,
	YamlKeyMarkxusPromptFormat:     EnvKeyMarkxusPromptFormat,
	YamlKeyNexusApiKey:             EnvKeyNexusApiKey,
	YamlKeyNexusUrlGetModFormat:    EnvKeyNexusUrlGetModFormat,
	YamlKeyMarkxusUrlModPageFormat: EnvKeyMarkxusUrlModPageFormat,
	YamlKeyMarkdownHeaderFormat:    EnvKeyMarkdownHeaderFormat,
	YamlKeyFallbackGameCode:        EnvKeyFallbackGameCode,
	YamlKeyInteractive:             EnvKeyInteractive,
}

var envToYaml = map[string]string{
	EnvKeyGenAiApiKey:             YamlKeyGenAiApiKey,
	EnvKeyGenAiModelName:          YamlKeyGenAiModelName,
	EnvKeyMarkxusPromptFormat:     YamlKeyMarkxusPromptFormat,
	EnvKeyNexusApiKey:             YamlKeyNexusApiKey,
	EnvKeyNexusUrlGetModFormat:    YamlKeyNexusUrlGetModFormat,
	EnvKeyMarkxusUrlModPageFormat: YamlKeyMarkxusUrlModPageFormat,
	EnvKeyMarkdownHeaderFormat:    YamlKeyMarkdownHeaderFormat,
	EnvKeyFallbackGameCode:        YamlKeyFallbackGameCode,
	EnvKeyInteractive:             YamlKeyInteractive,
}

func EnvToYaml(key string) string {
	return envToYaml[key]
}

func YamlToEnv(key string) string {
	return yamlToEnv[key]
}

type ConfigType string

var (
	ConfigTypeGlobal ConfigType = "global"
	ConfigTypeLocal  ConfigType = "local"
)
