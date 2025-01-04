package markxus

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/slainless/markxus/nexus"
)

type Generated struct {
	Mod     nexus.SchemaMod
	Content string
	Header  string
	Error   error
}

func (c *Markxus) Generate(
	ctx context.Context,
	gameCode string,
	modId string,
	options ...GenerationContextOption,
) (*Generated, error) {
	generationCtx := &GenerationContext{}

	for _, option := range options {
		option(generationCtx)
	}

	mod, err := c.nexus.GetMod(ctx, gameCode, modId)
	if err != nil {
		return nil, err
	}

	if generationCtx.OnModFetched != nil {
		err := generationCtx.OnModFetched(ctx, mod)
		if err != nil {
			return nil, err
		}
	}

	mod.MetadataPageUrl = fmt.Sprintf(c.options.UrlModPageFormat, gameCode, modId)

	if generationCtx.CategoryIconMap != nil {
		if icon := generationCtx.CategoryIconMap[mod.CategoryId]; icon != nil {
			mod.MetadataCategoryIcon = icon.Icon
		}
	}

	header, err := processTemplate(
		c.options.MarkdownHeaderTemplate,
		mod,
	)
	if err != nil {
		return nil, err
	}

	if generationCtx.OnHeaderCreated != nil {
		err := generationCtx.OnHeaderCreated(ctx, header)
		if err != nil {
			return nil, err
		}
	}

	prompt, err := processTemplate(
		c.options.LlmPromptTemplate,
		mod,
	)
	if err != nil {
		return nil, err
	}

	if generationCtx.OnPromptCreated != nil {
		err := generationCtx.OnPromptCreated(ctx, prompt)
		if err != nil {
			return nil, err
		}
	}

	output, err := c.llm.Send(ctx, prompt, mod, generationCtx.OnLlmStreamConsuming)
	if err != nil {
		return nil, err
	}

	return &Generated{
		Content: output,
		Mod:     *mod,
		Header:  header,
		Error:   err,
	}, nil
}

func processTemplate(format *template.Template, mod *nexus.SchemaMod) (string, error) {
	builder := &strings.Builder{}
	err := format.Execute(builder, mod)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
