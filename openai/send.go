package openai

import (
	"context"
	"errors"
	"io"

	"github.com/sashabaranov/go-openai"
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/nexus"
)

func (c *OpenAiClient) Send(
	ctx context.Context,
	prompt string,
	mod *nexus.SchemaMod,
	hook markxus.LlmStreamConsumeHook,
) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: c.options.ModelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
	stream, err := c.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	var output string
	err = processResponse(ctx, c.client, stream, req, &output, hook)
	if err != nil {
		return "", err
	}

	return output, nil
}

func processResponse(
	ctx context.Context,
	client *openai.Client,
	stream *openai.ChatCompletionStream,
	req openai.ChatCompletionRequest,
	output *string,
	hook markxus.LlmStreamConsumeHook,
) error {
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
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

		for _, choice := range response.Choices {
			*output += choice.Delta.Content

			switch choice.FinishReason {
			case openai.FinishReasonStop, openai.FinishReasonNull:
				continue
			case openai.FinishReasonLength:
				stream.Close()
				req.Messages = append(req.Messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: *output,
				}, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: "Continue where you left off and remember your previous task.",
				})
				stream, err := client.CreateChatCompletionStream(ctx, req)
				if err != nil {
					return err
				}

				return processResponse(ctx, client, stream, req, output, hook)
			case openai.FinishReasonContentFilter, openai.FinishReasonFunctionCall, openai.FinishReasonToolCalls:
				return &AIGenerationError{
					Reason: choice.FinishReason,
				}
			}
		}
	}

	return nil
}
