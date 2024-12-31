package internal

import (
	"os"

	"github.com/slainless/markxus/cli/markxus/config"
)

func OpenConfig(configType config.ConfigType) (*os.File, error) {
	return os.OpenFile(ConfigPath(configType), os.O_RDWR, 0600)
}
