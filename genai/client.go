package genai

import (
	"context"
	"errors"

	ai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GenAiClient struct {
	client  *ai.Client
	options *ClientOptions
}

func NewGenAiClient(ctx context.Context, options ...ClientOption) (*GenAiClient, error) {
	genaiOptions := &ClientOptions{
		ModelName: DefaultModelName,
	}

	for _, option := range options {
		option(genaiOptions)
	}

	if genaiOptions.ApiKey == "" {
		return nil, ErrNoApiKey
	}

	client, err := ai.NewClient(ctx, option.WithAPIKey(genaiOptions.ApiKey))
	if err != nil {
		panic(err)
	}

	return &GenAiClient{
		client:  client,
		options: genaiOptions,
	}, nil
}

var ErrNoApiKey = errors.New("no generative ai api key provided")
