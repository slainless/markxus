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
	KeyGenAiApiKey             = "genai_api_key"
	KeyGenAiModelName          = "genai_model_name"
	KeyMarkxusPromptFormat     = "genai_prompt_format"
	KeyNexusApiKey             = "nexus_api_key"
	KeyNexusUrlGetModFormat    = "nexus_url_get_mod_format"
	KeyMarkxusUrlModPageFormat = "nexus_url_mod_page_format"
	KeyMarkdownHeaderFormat    = "markdown_header_format"
	KeyFallbackGameCode        = "fallback_game_code"
)
