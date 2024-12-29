package nexus

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
)

func (c *Client) GetMod(ctx context.Context, gameCode string, modId string) (*SchemaMod, error) {
	raw, err := c.driver.Get(ctx, c.apiKey, fmt.Sprintf(c.urlGetModFormat, gameCode, modId))
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
