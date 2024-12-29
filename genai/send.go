package genai

import (
	"context"

	ai "github.com/google/generative-ai-go/genai"
	"github.com/slainless/markxus/nexus"
)

func (c *GenAiClient) Send(ctx context.Context, prompt string, mod *nexus.SchemaMod) (string, error) {
	model := c.client.GenerativeModel(c.options.ModelName)

	cs := model.StartChat()

	initial, err := cs.SendMessage(ctx, ai.Text(prompt))
	if err != nil {
		return "", err
	}

	var output string
	err = processResponse(initial, &output, ctx, cs)
	return output, err
}

func processResponse(response *ai.GenerateContentResponse, output *string, ctx context.Context, cs *ai.ChatSession) error {
	for _, candidate := range response.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				switch value := part.(type) {
				case ai.Text:
					*output += string(value)
				}
			}
		}

		switch candidate.FinishReason {
		case ai.FinishReasonUnspecified, ai.FinishReasonStop:
			return nil
		case ai.FinishReasonMaxTokens:
			response, err := cs.SendMessage(ctx, ai.Text("Continue"))
			if err != nil {
				return err
			}

			return processResponse(response, output, ctx, cs)
		case ai.FinishReasonOther, ai.FinishReasonRecitation, ai.FinishReasonSafety:
			return &AIGenerationError{
				Reason: candidate.FinishReason,
			}
		}
	}

	return nil
}
