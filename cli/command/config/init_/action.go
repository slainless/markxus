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
	configType := config.Config.Common.ConfigType.Selected()

	err := writeConfig(configType)
	if err != nil {
		return err
	}

	aiKey, nexusKey, err := PromptApiKey(ctx)
	if err != nil {
		return err
	}

	if aiKey != "" && nexusKey != "" {
		_ = config.SetKeyring(config.YamlKeyGenAiApiKey, aiKey)
		_ = config.SetKeyring(config.YamlKeyNexusApiKey, nexusKey)
	}

	theme := style.GetTheme()
	var cardContent strings.Builder
	fmt.Fprintf(&cardContent,
		lipgloss.NewStyle().
			Render("Config type - %s\nPath\n%s\n\nConfiguration has been initiated."),
		theme.Focused.Description.Render(string(configType)),
		theme.Focused.Card.Render(
			theme.Focused.Description.Render(configPath(configType)),
		),
	)

	fmt.Println(style.Card().Render(cardContent.String()))

	return err
}
