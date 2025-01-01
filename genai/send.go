package genai

import (
	"context"

	ai "github.com/google/generative-ai-go/genai"
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/nexus"
	"google.golang.org/api/iterator"
)

func (c *GenAiClient) Send(
	ctx context.Context,
	prompt string,
	mod *nexus.SchemaMod,
	hook markxus.LlmStreamConsumeHook,
) (string, error) {
	model := c.client.GenerativeModel(c.options.ModelName)

	cs := model.StartChat()
	iter := cs.SendMessageStream(ctx, ai.Text(prompt))

	var output string
	err := processResponse(iter, &output, ctx, cs, hook)
	if err != nil {
		return "", err
	}

	return output, err
}

func processResponse(
	iter *ai.GenerateContentResponseIterator,
	output *string,
	ctx context.Context,
	cs *ai.ChatSession,
	hook markxus.LlmStreamConsumeHook,
) error {
	shouldContinue := false
	for {
		response, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}

		if hook != nil {
			err := hook(ctx, response, output)
			if err != nil {
				return err
			}
		}

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
				continue
			case ai.FinishReasonMaxTokens:
				shouldContinue = true
				continue
			case ai.FinishReasonOther, ai.FinishReasonRecitation, ai.FinishReasonSafety:
				return &AIGenerationError{
					Reason: candidate.FinishReason,
				}
			}
		}
	}

	if shouldContinue {
		iter := cs.SendMessageStream(ctx, ai.Text("Continue where you left off and remember your previous task."))
		return processResponse(iter, output, ctx, cs, hook)
	}

	return nil
}
