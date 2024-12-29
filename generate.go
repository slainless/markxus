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

func (c *Markxus) Generate(ctx context.Context, gameCode string, modId string) (*Generated, error) {
	mod, err := c.nexus.GetMod(ctx, gameCode, modId)
	if err != nil {
		return nil, err
	}

	mod.PageUrl = fmt.Sprintf(c.options.UrlModPageFormat, gameCode, modId)

	header, err := processHeader(
		c.options.MarkdownHeaderTemplate,
		mod,
	)
	if err != nil {
		return nil, err
	}

	prompt := fmt.Sprintf(c.options.GenAiPromptFormat, mod.Description)
	output, err := c.llm.Send(ctx, prompt, mod)
	if len(output) < 1 {
		return nil, err
	}

	return &Generated{
		Content: output,
		Mod:     *mod,
		Header:  header,
		Error:   err,
	}, nil
}

func processHeader(format *template.Template, mod *nexus.SchemaMod) (string, error) {
	builder := &strings.Builder{}
	err := format.Execute(builder, mod)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
