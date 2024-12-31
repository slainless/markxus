package set

import (
	"fmt"
	"io"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/slainless/markxus/cli/markxus/command/config/internal"
	"github.com/slainless/markxus/cli/markxus/config"
)

func readConfig(configType config.ConfigType) (*os.File, map[string]any, error) {
	configPath := config.ConfigPath(configType)
	if configPath == "" {
		return nil, nil, fmt.Errorf("invalid file path")
	}

	handle, err := internal.OpenConfig(configType)
	if err != nil {
		return nil, nil, err
	}

	raw, err := io.ReadAll(handle)
	if err != nil {
		return nil, nil, err
	}

	var data map[string]any
	err = yaml.Unmarshal(raw, &data)
	if err != nil {
		return nil, nil, err
	}

	return handle, data, nil
}
