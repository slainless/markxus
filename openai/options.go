package openai

import "github.com/sashabaranov/go-openai"

type ClientOptions struct {
	ApiKey    string
	ModelName string
}

const DefaultModelName = openai.GPT4oMini20240718

type ClientOption func(*ClientOptions)

func WithApiKey(key string) ClientOption {
	return func(co *ClientOptions) {
		co.ApiKey = key
	}
}

func WithModelName(name string) ClientOption {
	return func(co *ClientOptions) {
		co.ModelName = name
	}
}
