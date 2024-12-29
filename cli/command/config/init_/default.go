package init_

import (
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
)

func defaultConfig() map[string]any {
	kv := map[string]any{}

	kv[config.YamlKeyGenAiApiKey] = ""
	kv[config.YamlKeyGenAiModelName] = genai.DefaultModelName
	kv[config.YamlKeyMarkxusPromptFormat] = markxus.DefaultGenAiPromptFormat

	kv[config.YamlKeyNexusApiKey] = ""
	kv[config.YamlKeyNexusUrlGetModFormat] = nexus.DefaultUrlGetModFormat
	kv[config.YamlKeyMarkxusUrlModPageFormat] = markxus.DefaultUrlModPageFormat
	kv[config.YamlKeyMarkdownHeaderFormat] = markxus.DefaultMarkdownHeaderFormat

	return kv
}
