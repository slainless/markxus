package cmd_init

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/slainless/markxus/cli/config"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

func action(ctx context.Context, c *cli.Command) error {
	if initConfig.force {
		return writeToFile()
	}

	configType := "global"
	typeForm := huh.NewForm(
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

	err := typeForm.RunWithContext(ctx)
	if err != nil {
		return err
	}

	handle, err := os.OpenFile(configPath(), os.O_WRONLY, 0600)
	if os.IsNotExist(err) {
		return writeToFile()
	} else if err != nil {
		return err
	}
	defer handle.Close()

	var confirm bool
	confirmForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Config already exist, overwrite?").
				Value(&confirm),
		),
	)

	err = confirmForm.RunWithContext(ctx)
	if err != nil {
		return err
	}

	if !confirm {
		return nil
	}

	err = handle.Truncate(0)
	if err != nil {
		return err
	}

	data, err := createConfigData()
	if err != nil {
		return err
	}

	_, err = handle.Write(data)
	return err
}

func configPath() string {
	switch initConfig.configType.String() {
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

func writeToFile() error {
	data, err := createConfigData()
	if err != nil {
		return err
	}

	path := configPath()
	if path == "" {
		return fmt.Errorf("invalid config type: %s", initConfig.configType.String())
	}

	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
