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
}

var Config MarkxusConfig
