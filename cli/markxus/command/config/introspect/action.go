package introspect

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	theme := style.GetTheme()

	configType := config.ConfigType(c.Args().First())
	configPath := config.ConfigPath(configType)

	var header strings.Builder

	header.WriteString(
		theme.Focused.Title.Render("Configuration Path") + "\n\n",
	)

	if configPath == "" || configType == config.ConfigTypeGlobal {
		header.WriteString(createConfigEntry(config.ConfigTypeGlobal, theme))

		if config.YamlSourceGlobalError != nil {
			header.WriteString("\n" + theme.Focused.ErrorMessage.Render(config.YamlSourceGlobalError.Error()))
		}
	}

	if configPath == "" {
		header.WriteString("\n")
	}

	if configPath == "" || configType == config.ConfigTypeLocal {
		header.WriteString(createConfigEntry(config.ConfigTypeLocal, theme))

		if config.YamlSourceGlobalError != nil {
			header.WriteString("\n" + theme.Focused.ErrorMessage.Render(config.YamlSourceGlobalError.Error()))
		}
	}

	var vars strings.Builder

	vars.WriteString(theme.Focused.Title.Render("Variables") + "\n")

	if configType == config.ConfigTypeGlobal {
		vars.WriteString(theme.Focused.Description.Render("Resolved from global config") + "\n\n")
	} else if configType == config.ConfigTypeLocal {
		vars.WriteString(theme.Focused.Description.Render("Resolved from local config") + "\n\n")
	} else {
		vars.WriteString(theme.Focused.Description.Render("Resolved from all sources") + "\n\n")
	}

	vars.WriteString(createVarEntry(configType, config.YamlKeyLlmProvider, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyLlmApiKey, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyLlmModelName, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyMarkxusPromptFormat, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyNexusApiKey, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyNexusUrlGetModFormat, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyNexusUrlGetFilesFormat, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyMarkxusUrlModPageFormat, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyMarkdownHeaderFormat, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyOutputDir, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyOverwriteOutput, theme) + "\n\n")
	vars.WriteString(createVarEntry(configType, config.YamlKeyFallbackGameCode, theme))

	fmt.Println(style.Card().Render(header.String()))
	fmt.Println(style.Card().Render(vars.String()))

	return nil
}

func createVarEntry(configType config.ConfigType, yamlKey string, theme *huh.Theme) string {
	var value any
	if configType == "" {
		value = config.Resolve(config.YamlToEnv(yamlKey))
	} else {
		value = config.ResolveFromYaml(configType, yamlKey)
	}

	var v string
	if value == nil || value == "" {
		v = theme.Blurred.TextInput.Placeholder.Render("(empty)")
	} else {
		v = theme.Focused.Description.Render(fmt.Sprint(value))
	}

	return fmt.Sprintf(
		"%s\n%s",
		config.YamlToEnv(yamlKey),
		v,
	)
}

func createConfigEntry(configType config.ConfigType, theme *huh.Theme) string {
	var title string
	switch configType {
	case config.ConfigTypeGlobal:
		title = "Global"
	case config.ConfigTypeLocal:
		title = "Local"
	}

	return lipgloss.NewStyle().
		Render(
			fmt.Sprintf(
				"%s\n%s",
				title,
				theme.Focused.Description.Render(config.ConfigPath(configType)),
			),
		)
}
