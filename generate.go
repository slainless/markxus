package markxus

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/google/generative-ai-go/genai"
	"github.com/slainless/markxus/nexus"
)

type Generated struct {
	Mod     nexus.SchemaMod
	Content string
	Header  string
}

func (c *Markxus) Generate(ctx context.Context, gameCode string, modId string) (*Generated, error) {
	mod, err := c.nexus.GetMod(ctx, gameCode, modId)
	if err != nil {
		return nil, err
	}

	mod.PageUrl = fmt.Sprintf(c.genaiOptions.urlModPageFormat, gameCode, modId)

	header, err := processHeader(
		c.genaiOptions.markdownHeaderTemplate,
		mod,
	)
	if err != nil {
		return nil, err
	}

	model := c.genai.GenerativeModel(c.genaiOptions.genAiModelName)

	cs := model.StartChat()

	initial, err := cs.SendMessage(ctx, genai.Text(fmt.Sprintf(c.genaiOptions.genAiPromptFormat, mod.Description)))
	if err != nil {
		return nil, err
	}

	var output string
	err = processResponse(initial, &output, ctx, cs)

	return &Generated{
		Content: output,
		Mod:     *mod,
		Header:  header,
	}, err
}

func processResponse(response *genai.GenerateContentResponse, output *string, ctx context.Context, cs *genai.ChatSession) error {
	for _, candidate := range response.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				switch value := part.(type) {
				case genai.Text:
					*output += string(value)
				}
			}
		}

		switch candidate.FinishReason {
		case genai.FinishReasonUnspecified, genai.FinishReasonStop:
			return nil
		case genai.FinishReasonMaxTokens:
			response, err := cs.SendMessage(ctx, genai.Text("Continue"))
			if err != nil {
				return err
			}

			return processResponse(response, output, ctx, cs)
		case genai.FinishReasonOther, genai.FinishReasonRecitation, genai.FinishReasonSafety:
			return &AIGenerationError{
				Reason: candidate.FinishReason,
			}
		}
	}

	return nil
}

func processHeader(format *template.Template, mod *nexus.SchemaMod) (string, error) {
	builder := &strings.Builder{}
	err := format.Execute(builder, mod)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
