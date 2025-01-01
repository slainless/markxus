package openai

import "github.com/sashabaranov/go-openai"

type OpenAiClient struct {
	client  *openai.Client
	options *ClientOptions
}

func NewOpenAiClient(options ...ClientOption) (*OpenAiClient, error) {
	openaiOptions := &ClientOptions{
		ModelName: DefaultModelName,
	}

	for _, option := range options {
		option(openaiOptions)
	}

	client := openai.NewClient(openaiOptions.ApiKey)

	return &OpenAiClient{
		options: openaiOptions,
		client:  client,
	}, nil
}
