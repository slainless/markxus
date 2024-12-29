package genai

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

type AIGenerationError struct {
	Reason genai.FinishReason
}

func (e *AIGenerationError) Error() string {
	return fmt.Sprintf("AI Generation stopped, caused by: %v", formatReason(e.Reason))
}

func formatReason(reason genai.FinishReason) string {
	switch reason {
	case genai.FinishReasonUnspecified:
		return "Unspecified"
	case genai.FinishReasonStop:
		return "Natural stop"
	case genai.FinishReasonMaxTokens:
		return "Maximum token output reached"
	case genai.FinishReasonSafety:
		return "Safety reasons"
	case genai.FinishReasonRecitation:
		return "Recitation reasons"
	case genai.FinishReasonOther:
		return "Other"
	}

	return reason.String()
}
