package resty

import (
	"context"

	r "github.com/go-resty/resty/v2"
)

type RestyClient struct {
	Client *r.Client
}

func (c *RestyClient) Get(ctx context.Context, apiKey string, url string) (string, error) {
	res, err := c.Client.R().SetHeader("ApiKey", apiKey).Get(url)
	if err != nil {
		return "", err
	}

	status := res.StatusCode()
	if 200 <= status && status < 300 {
		return "", &RestyError{
			Response: res,
		}
	}

	return res.String(), nil
}
