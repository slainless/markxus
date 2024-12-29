package init_

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/urfave/cli/v3"
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

	_ = config.SetKeyring(config.YamlKeyGenAiApiKey, aiKey)
	_ = config.SetKeyring(config.YamlKeyNexusApiKey, nexusKey)

	theme := style.GetTheme()
	var cardContent strings.Builder
	fmt.Fprintf(&cardContent,
		lipgloss.NewStyle().
			Render("Config type - %s\nPath\n%s\n\nConfiguration has been initiated."),
		theme.Focused.Description.Render(configType),
		theme.Focused.Card.Render(
			theme.Focused.Description.Render(configPath(configType)),
		),
	)

	fmt.Println(style.Card().Render(cardContent.String()))

	return err
}
