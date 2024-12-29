package nexus

import (
	"errors"
)

type Client struct {
	options *ClientOptions
}

func NewClient(options ...ClientOption) (*Client, error) {
	clientOptions := &ClientOptions{
		UrlGetModFormat: DefaultUrlGetModFormat,
	}

	for _, option := range options {
		option(clientOptions)
	}

	if clientOptions.ApiKey == "" {
		return nil, ErrorNoApiKey
	}

	if clientOptions.Driver == nil {
		return nil, ErrorNoDriver
	}

	client := &Client{
		options: clientOptions,
	}

	return client, nil
}

var ErrorNoApiKey = errors.New("no nexus api key provided")
var ErrorNoDriver = errors.New("no nexus http driver provided")
