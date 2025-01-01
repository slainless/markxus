package init_

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/command/config/internal"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	configType := internal.ConfigType()

	err := internal.WriteFileWithCheck(configType, config.CreateDefaultConfig())
	if err != nil {
		return err
	}

	aiKey, nexusKey, err := PromptApiKey(ctx)
	if err != nil {
		return err
	}

	if aiKey != "" && nexusKey != "" {
		err := config.SetKeyring(config.YamlKeyLlmApiKey, aiKey)

		var sb strings.Builder
		sb.WriteString("Config initialized, but failed to set keyring:\n")
		if err != nil {
			sb.WriteString(err.Error())
		}

		err = config.SetKeyring(config.YamlKeyNexusApiKey, nexusKey)
		if err != nil {
			sb.WriteString(err.Error())
		}

		return errors.New(sb.String())
	}

	theme := style.GetTheme()
	var cardContent strings.Builder
	fmt.Fprintf(&cardContent,
		lipgloss.NewStyle().
			Render("Config type - %s\nPath\n%s\n\nConfiguration has been initialized."),
		theme.Focused.Description.Render(string(configType)),
		theme.Focused.Description.Render(config.ConfigPath(configType)),
	)

	fmt.Println(style.Card().Render(cardContent.String()))

	return err
}
