package internal

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/slainless/markxus/cli/markxus/config"
)

func WriteFileWithCheck(configType config.ConfigType, data map[string]any) error {
	configPath := ConfigPath(configType)
	if configPath == "" {
		return fmt.Errorf("invalid file path")
	}

	handle, err := OpenConfig(configType)
	if os.IsNotExist(err) {
		return WriteToFile(ConfigPath(configType), data)
	} else if err != nil {
		return err
	}

	defer handle.Close()

	if !IsOverwrite() {
		return fmt.Errorf("config already exists!\nuse --overwrite to overwrite it")
	}

	return WriteToHandle(handle, data)
}

func WriteToHandle(handle *os.File, data map[string]any) error {
	err := handle.Truncate(0)
	if err != nil {
		return err
	}

	raw, err := Marshal(data)
	if err != nil {
		return err
	}

	_, err = handle.WriteAt(raw, 0)
	return err
}

func WriteToFile(path string, data map[string]any) error {
	raw, err := Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, raw, 0600)
	if err != nil {
		return err
	}

	return nil
}

func Marshal(data map[string]any) ([]byte, error) {
	return yaml.MarshalWithOptions(data, yaml.UseLiteralStyleIfMultiline(true))
}
