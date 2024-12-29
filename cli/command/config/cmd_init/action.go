package cmd_init

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

func action(ctx context.Context, c *cli.Command) error {
	if initConfig.force {
		return writeToFile(configPath(initConfig.configType.String()))
	}

	configType, err := PromptConfigType(ctx)
	if err != nil {
		return err
	}

	err = writeConfig(ctx, configType)
	if err != nil {
		return err
	}

	aiKey, nexusKey, err := PromptApiKey(ctx)
	if err != nil {
		return err
	}

	_ = config.SetKeyring(config.KeyGenAiApiKey, aiKey)
	_ = config.SetKeyring(config.KeyNexusApiKey, nexusKey)

	theme := internal.GetTheme()
	var cardContent strings.Builder
	fmt.Fprintf(&cardContent,
		lipgloss.NewStyle().
			Render("Config type - %s\nPath\n%s\n\nConfiguration has been initiated."),
		theme.Focused.Description.Render(configType),
		theme.Focused.Card.Render(
			theme.Focused.Description.Render(configPath(configType)),
		),
	)

	fmt.Println(
		lipgloss.NewStyle().
			Padding(1, 2).
			Width(40).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.FieldSeparator.GetForeground()).
			Render(cardContent.String()),
	)

	return err
}

func configPath(configType string) string {
	switch configType {
	case "global":
		return config.GlobalConfigPath
	case "local":
		return config.LocalConfigPath
	}

	return ""
}

func createConfigData() ([]byte, error) {
	data, err := yaml.Marshal(defaultConfig())
	if err != nil {
		return nil, err
	}

	return data, nil
}
