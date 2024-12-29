package cmd_init

import (
	"context"
	"os"
)

func writeConfig(ctx context.Context, configType string) error {
	handle, err := os.OpenFile(configPath(configType), os.O_WRONLY, 0600)
	if os.IsNotExist(err) {
		return writeToFile(configPath(configType))
	} else if err != nil {
		return err
	}

	defer handle.Close()

	overwrite, err := PromptFileOverwrite(ctx)
	if err != nil {
		return err
	}

	if !overwrite {
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
