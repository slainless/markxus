package init_

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/slainless/markxus/cli/markxus/config"
)

func writeConfig(configType config.ConfigType) error {
	handle, err := os.OpenFile(configPath(configType), os.O_WRONLY, 0600)
	if os.IsNotExist(err) {
		return writeToFile(configPath(configType))
	} else if err != nil {
		return err
	}

	defer handle.Close()

	if !config.Config.Common.Overwrite {
		return fmt.Errorf("config already exists!\nuse --overwrite to overwrite it")
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

func writeToFile(path string) error {
	data, err := createConfigData()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func configPath(configType config.ConfigType) string {
	switch configType {
	case config.ConfigTypeGlobal:
		return config.GlobalConfigPath
	case config.ConfigTypeLocal:
		return config.LocalConfigPath
	}

	return ""
}

func createConfigData() ([]byte, error) {
	data, err := yaml.MarshalWithOptions(defaultConfig(), yaml.UseLiteralStyleIfMultiline(true))
	if err != nil {
		return nil, err
	}

	return data, nil
}
