package nexus

import (
	"context"
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
)

func (c *Client) GetMod(ctx context.Context, gameCode string, modId string) (*SchemaMod, error) {
	derivedCtx, cancel := context.WithCancel(ctx)

	var modErr, filesErr error
	wg := sync.WaitGroup{}
	wg.Add(2)

	var mod SchemaMod
	var files SchemaModFiles

	go func() {
		defer wg.Done()

		var raw string
		raw, modErr = c.options.
			Driver.Get(
			derivedCtx,
			c.options.ApiKey,
			fmt.Sprintf(c.options.UrlGetModFormat, gameCode, modId),
		)

		if modErr != nil {
			cancel()
			return
		}

		modErr = sonic.UnmarshalString(raw, &mod)
		if modErr != nil {
			cancel()
			return
		}
	}()

	go func() {
		defer wg.Done()

		var raw string
		raw, filesErr = c.options.
			Driver.Get(
			derivedCtx,
			c.options.ApiKey,
			fmt.Sprintf(c.options.UrlGetFilesFormat, gameCode, modId),
		)

		if filesErr != nil {
			cancel()
			return
		}

		filesErr = sonic.UnmarshalString(raw, &files)
		if filesErr != nil {
			cancel()
			return
		}
	}()

	wg.Wait()

	if modErr != nil && modErr != derivedCtx.Err() {
		return nil, modErr
	}

	if filesErr != nil && filesErr != derivedCtx.Err() {
		return nil, filesErr
	}

	if derivedCtx.Err() != nil {
		return nil, derivedCtx.Err()
	}

	mod.Files = &files
	return &mod, nil
}
