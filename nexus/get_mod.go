package nexus

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
)

func (c *Client) GetMod(ctx context.Context, gameCode string, modId string) (*SchemaMod, error) {
	raw, err := c.options.
		Driver.Get(
		ctx,
		c.options.ApiKey,
		fmt.Sprintf(c.options.UrlGetModFormat, gameCode, modId),
	)

	if err != nil {
		return nil, err
	}

	var mod SchemaMod
	err = sonic.UnmarshalString(raw, &mod)
	if err != nil {
		return nil, err
	}

	return &mod, nil
}
