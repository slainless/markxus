package nexus

import (
	"errors"
)

type Client struct {
	driver HttpClient
	apiKey string

	urlGetModFormat string
}

func NewClient(options ...ClientOption) (*Client, error) {
	client := &Client{
		urlGetModFormat: DefaultUrlGetModFormat,
	}
	for _, option := range options {
		option(client)
	}

	if client.apiKey == "" {
		return nil, ErrorNoApiKey
	}

	if client.driver == nil {
		return nil, ErrorNoDriver
	}

	return client, nil
}

var ErrorNoApiKey = errors.New("no nexus api key provided")
var ErrorNoDriver = errors.New("no nexus http driver provided")
