package resty

import (
	"context"

	r "github.com/go-resty/resty/v2"
)

type RestyClient struct {
	Client *r.Client
}

func NewRestyClient(base ...*r.Client) *RestyClient {
	var c *r.Client
	if len(base) == 1 && base[0] != nil {
		c = base[0]
	} else {
		c = r.New()
	}
	return &RestyClient{Client: c}
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
