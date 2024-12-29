package resty

import (
	r "github.com/go-resty/resty/v2"
	"github.com/slainless/markxus/nexus"
)

func WithRestyClient(client ...*r.Client) nexus.ClientOption {
	var c *r.Client
	if len(client) == 1 && client[0] != nil {
		c = client[0]
	}
	return func(co *nexus.ClientOptions) {
		co.Driver = &RestyClient{Client: c}
	}
}
