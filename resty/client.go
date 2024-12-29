package resty

import (
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
