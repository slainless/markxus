package cmd_init

import (
	"context"

	"github.com/charmbracelet/huh"
)

func PromptApiKey(ctx context.Context) (ai string, nexus string, err error) {

	var aiKey string
	var nexusKey string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Nexus Mods API Key").
				Value(&nexusKey),
			huh.NewInput().
				Title("Google Generative AI API Key").
				Value(&aiKey),
		),
	)

	err = form.RunWithContext(ctx)
	if err != nil {
		return "", "", err
	}

	return aiKey, nexusKey, nil
}
