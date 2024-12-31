package init_

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func PromptApiKey(ctx context.Context) (ai string, nexus string, err error) {
	var aiKey string
	var nexusKey string

	var form *huh.Form

	confirm := false
	form = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to set up API keys to OS keyring?").
				Description("Storing it in OS is more secure than storing it in config file.").
				WithButtonAlignment(lipgloss.Left).
				Value(&confirm),
		),
	)

	err = form.RunWithContext(ctx)
	if err != nil {
		return "", "", err
	}

	if !confirm {
		return "", "", nil
	}

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Configure API Keys").
				Description("These keys will be stored in OS keyring.\nYou can also set them via *markxus config set* command."),
			huh.NewInput().
				Title("Nexus Mods").
				Value(&nexusKey).
				Placeholder("API Key").
				Validate(func(value string) error {
					if value == "" {
						return fmt.Errorf("nexus mods api key is required")
					}

					return nil
				}),
			huh.NewInput().
				Title("Google Generative AI").
				Value(&aiKey).
				Placeholder("API Key").
				Validate(func(value string) error {
					if value == "" {
						return fmt.Errorf("google generative ai api key is required")
					}

					return nil
				}),
		),
	)

	err = form.RunWithContext(ctx)
	if err != nil {
		return "", "", err
	}

	return aiKey, nexusKey, nil
}
