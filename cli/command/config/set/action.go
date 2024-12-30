package set

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/command/config/internal"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	yamlKey, value, err := validateKv(c)
	if err != nil {
		return err
	}

	var previousValue string
	useKeyring := yamlKey == config.YamlKeyGenAiApiKey || yamlKey == config.YamlKeyNexusApiKey
	if useKeyring {
		err = config.SetKeyring(yamlKey, value)
		if err != nil {
			return err
		}
	} else {
		configType := internal.ConfigType()

		handle, data, err := readConfig(configType)
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found.\nInit first with *markxus config init -t %s*", configType)
		} else if err != nil {
			return err
		}
		defer handle.Close()

		if data[yamlKey] != nil {
			previousValue = fmt.Sprint(data[yamlKey])
		}

		data[yamlKey] = value
		err = internal.WriteToHandle(handle, data)
		if err != nil {
			return err
		}
	}

	var displayConfigType string
	if useKeyring {
		displayConfigType = "keyring"
	} else {
		displayConfigType = string(internal.ConfigType())
	}

	theme := style.GetTheme()
	var cardContent strings.Builder
	fmt.Fprintf(&cardContent,
		lipgloss.NewStyle().Render("Config type - %s\nKey - %s\n\n"),
		theme.Focused.Description.Render(displayConfigType),
		theme.Focused.Card.Render(theme.Focused.Description.Render(yamlKey)),
	)

	if previousValue != "" {
		fmt.Fprintf(&cardContent,
			lipgloss.NewStyle().Render("Previous value\n%s\n"),
			theme.Focused.Card.Render(theme.Focused.Description.Render(previousValue)),
		)
	}

	fmt.Fprintf(&cardContent,
		lipgloss.NewStyle().Render("New value\n%s\n\n"),
		theme.Focused.Card.Render(theme.Focused.Description.Render(value)),
	)
	cardContent.WriteString("Key has been set.")

	fmt.Println(style.Card().Render(cardContent.String()))
	return nil
}

func validateKv(c *cli.Command) (string, string, error) {
	key := c.Args().First()
	value := c.Args().Get(1)

	if key == "" {
		return "", "", fmt.Errorf("no key given")
	}

	if value == "" {
		return "", "", fmt.Errorf("no value given")
	}

	envKey := config.YamlToEnv(key)
	yamlKey := config.EnvToYaml(key)

	if envKey == "" && yamlKey == "" {
		return "", "", fmt.Errorf("key is not valid")
	}

	if yamlKey == "" {
		return key, value, nil
	}

	return yamlKey, value, nil
}
