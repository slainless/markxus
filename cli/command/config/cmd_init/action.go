package cmd_init

import (
	"context"

	"github.com/slainless/markxus/cli/markxus/config"
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
