package config

import (
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
)

func CreateDefaultConfig() map[string]any {
	kv := map[string]any{}

	kv[YamlKeyGenAiApiKey] = ""
	kv[YamlKeyGenAiModelName] = genai.DefaultModelName
	kv[YamlKeyMarkxusPromptFormat] = markxus.DefaultGenAiPromptFormat

	kv[YamlKeyNexusApiKey] = ""
	kv[YamlKeyNexusUrlGetModFormat] = nexus.DefaultUrlGetModFormat
	kv[YamlKeyMarkxusUrlModPageFormat] = markxus.DefaultUrlModPageFormat
	kv[YamlKeyMarkdownHeaderFormat] = markxus.DefaultMarkdownHeaderFormat
	kv[YamlKeyOutputDir] = FlagOutputDir.Value
	kv[YamlKeyOverwriteOutput] = FlagOverwriteOutput.Value
	kv[YamlKeyFallbackGameCode] = ""

	return kv
}
