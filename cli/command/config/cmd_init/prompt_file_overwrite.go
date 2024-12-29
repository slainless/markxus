package cmd_init

import (
	"context"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func PromptFileOverwrite(ctx context.Context) (bool, error) {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Config already exist").
				Description("Overwrite?").
				WithButtonAlignment(lipgloss.Left).
				Value(&confirm),
		),
	)

	err := form.RunWithContext(ctx)
	if err != nil {
		return false, err
	}

	return confirm, nil
}
