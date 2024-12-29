package nexus

import (
	"context"
)

type HttpClient interface {
	Get(ctx context.Context, apiKey string, url string) (string, error)
}
