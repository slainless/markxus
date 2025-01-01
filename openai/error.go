package openai

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type AIGenerationError struct {
	Reason openai.FinishReason
}

func (e *AIGenerationError) Error() string {
	return fmt.Sprintf("AI Generation stopped, caused by: %v", formatReason(e.Reason))
}

func formatReason(reason openai.FinishReason) string {
	switch reason {
	case openai.FinishReasonNull:
		return "Unspecified"
	case openai.FinishReasonStop:
		return "Natural stop"
	case openai.FinishReasonLength:
		return "Maximum token output reached"
	case openai.FinishReasonContentFilter:
		return "Safety reasons"
	case openai.FinishReasonToolCalls:
		return "Tool calls"
	case openai.FinishReasonFunctionCall:
		return "Function call"
	}

	return ""
}
