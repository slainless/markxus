package resty

import "context"

func (c *RestyClient) Get(ctx context.Context, apiKey string, url string) (string, error) {
	res, err := c.Client.R().SetHeader("ApiKey", apiKey).Get(url)
	if err != nil {
		return "", err
	}

	status := res.StatusCode()
	if 200 <= status && status < 300 {
		return res.String(), nil
	}

	return "", &RestyError{
		Response: res,
	}
}
