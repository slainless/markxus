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
	}
}

var Config MarkxusConfig

var (
	YamlKeyGenAiApiKey             = "genai_api_key"
	YamlKeyGenAiModelName          = "genai_model_name"
	YamlKeyMarkxusPromptFormat     = "genai_prompt_format"
	YamlKeyNexusApiKey             = "nexus_api_key"
	YamlKeyNexusUrlGetModFormat    = "nexus_url_get_mod_format"
	YamlKeyMarkxusUrlModPageFormat = "nexus_url_mod_page_format"
	YamlKeyMarkdownHeaderFormat    = "markdown_header_format"
	YamlKeyFallbackGameCode        = "fallback_game_code"
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
)
