package init_

import (
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
)

func defaultConfig() map[string]any {
	kv := map[string]any{}

	kv[config.KeyGenAiApiKey] = ""
	kv[config.KeyGenAiModelName] = genai.DefaultModelName
	kv[config.KeyMarkxusPromptFormat] = markxus.DefaultGenAiPromptFormat

	kv[config.KeyNexusApiKey] = ""
	kv[config.KeyNexusUrlGetModFormat] = nexus.DefaultUrlGetModFormat
	kv[config.KeyMarkxusUrlModPageFormat] = markxus.DefaultUrlModPageFormat
	kv[config.KeyMarkdownHeaderFormat] = markxus.DefaultMarkdownHeaderFormat

	return kv
}
