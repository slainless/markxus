package init_

import (
	"context"

	"github.com/charmbracelet/huh"
)

func PromptConfigType(ctx context.Context) (string, error) {
	configType := "global"
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select config type").
				Options(
					huh.NewOption("Global", "global"),
					huh.NewOption("Local", "local"),
				).
				Value(&configType),
		),
	)

	err := form.RunWithContext(ctx)
	if err != nil {
		return "", err
	}

	return configType, nil
}
