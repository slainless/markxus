package genai

type ClientOptions struct {
	ApiKey    string
	ModelName string
}

type ClientOption func(*ClientOptions)

const DefaultModelName = "gemini-1.5-flash"

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
