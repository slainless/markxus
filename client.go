package markxus

import (
	"context"
	_ "embed"
	"errors"

	"github.com/google/generative-ai-go/genai"
	"github.com/slainless/markxus/nexus"
	"google.golang.org/api/option"
)

type Markxus struct {
	genai        *genai.Client
	genaiOptions *markxusOptions

	nexus *nexus.Client
}

func NewMarkxus(ctx context.Context, nexusClient *nexus.Client, options ...MarkxusOption) (*Markxus, error) {
	var err error

	genaiOptions := &markxusOptions{
		genAiPromptFormat:      DefaultGenAiPromptFormat,
		genAiModelName:         DefaultGenAiModelName,
		urlModPageFormat:       DefaultUrlModPageFormat,
		markdownHeaderTemplate: DefaultMarkdownHeaderTemplate,
	}

	for _, option := range options {
		option(genaiOptions)
	}

	if genaiOptions.genAiApiKey == "" {
		return nil, ErrorGenAiNoApiKey
	}

	genai, err := genai.NewClient(ctx, option.WithAPIKey(genaiOptions.genAiApiKey))
	if err != nil {
		return nil, err
	}

	return &Markxus{
		nexus: nexusClient,
		genai: genai,
	}, nil
}

var ErrorGenAiNoApiKey = errors.New("no generative ai api key provided")
