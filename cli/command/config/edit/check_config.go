package edit

import (
	"errors"
	"os"

	"github.com/slainless/markxus/cli/markxus/config"
)

func isConfigExist(configType config.ConfigType) (bool, error) {
	path := config.ConfigPath(configType)
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return false, &os.PathError{
			Op:   "stat",
			Path: path,
			Err:  errors.New("path is a directory"),
		}
	}

	return true, nil
}
