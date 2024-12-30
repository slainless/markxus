package markxus

import (
	"context"

	"github.com/slainless/markxus/nexus"
)

type OnHeaderCreationHook func(ctx context.Context, header string) error
type OnModFetchedHook func(ctx context.Context, mod *nexus.SchemaMod) error

type GenerationContext struct {
	// sorted by call sequence
	OnModFetched         OnModFetchedHook
	OnHeaderCreated      OnHeaderCreationHook
	OnLlmStreamConsuming LlmStreamConsumeHook
}

type GenerationContextOption func(ctx *GenerationContext)

func WithOnHeaderCreation(f OnHeaderCreationHook) GenerationContextOption {
	return func(ctx *GenerationContext) {
		ctx.OnHeaderCreated = f
	}
}

func WithOnModFetched(f OnModFetchedHook) GenerationContextOption {
	return func(ctx *GenerationContext) {
		ctx.OnModFetched = f
	}
}

func WithOnLlmStreamConsuming(f LlmStreamConsumeHook) GenerationContextOption {
	return func(ctx *GenerationContext) {
		ctx.OnLlmStreamConsuming = f
	}
}
