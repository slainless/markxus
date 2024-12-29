package markxus

import (
	"context"

	"github.com/slainless/markxus/nexus"
)

type LlmClient interface {
	// If possible, [[LlmClient.Send]] must handle all common AI errors.
	// Only when it is not possible to handle such errors that
	// this method should return. Returning a non-empty string with error
	// will result in [[Generated]] with non-empty error.
	//
	// Some of those common error such as:
	//   - Max token limit
	//   - Safety error
	//   - Recitation error
	//
	// It is highly recommended to create a chat to allow resuming
	// in case max token limit hit. In this case, the Gen AI client
	// must manually send "continue" prompt and append the next response
	// batch to the result.
	Send(ctx context.Context, prompt string, mod *nexus.SchemaMod) (string, error)
}
