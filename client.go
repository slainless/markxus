package markxus

import (
	_ "embed"

	"github.com/slainless/markxus/nexus"
)

type Markxus struct {
	options *MarkxusOptions

	llm   LlmClient
	nexus *nexus.Client
}

func NewMarkxus(nexusClient *nexus.Client, genAiClient LlmClient, options ...MarkxusOption) *Markxus {
	markxusOptions := &MarkxusOptions{
		GenAiPromptFormat:      DefaultGenAiPromptFormat,
		UrlModPageFormat:       DefaultUrlModPageFormat,
		MarkdownHeaderTemplate: DefaultMarkdownHeaderTemplate,
	}

	for _, option := range options {
		option(markxusOptions)
	}

	return &Markxus{
		nexus:   nexusClient,
		llm:     genAiClient,
		options: markxusOptions,
	}
}
