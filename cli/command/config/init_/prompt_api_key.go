package init_

import (
	"context"

	"github.com/charmbracelet/huh"
)

func PromptApiKey(ctx context.Context) (ai string, nexus string, err error) {

	var aiKey string
	var nexusKey string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Description("These keys will be stored in OS keyring.\nYou can also set them via *markxus config set* command."),
			huh.NewInput().Title("Nexus Mods").Value(&nexusKey).Placeholder("API Key"),
			huh.NewInput().Title("Google Generative AI").Value(&aiKey).Placeholder("API Key"),
		),
	)

	err = form.RunWithContext(ctx)
	if err != nil {
		return "", "", err
	}

	return aiKey, nexusKey, nil
}
