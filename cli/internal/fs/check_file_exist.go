package fs

import (
	"errors"
	"os"

	"github.com/slainless/markxus/cli/markxus/config"
)

func IsFileExist(configType config.ConfigType) (os.FileInfo, bool, error) {
	path := config.ConfigPath(configType)
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	if stat.IsDir() {
		return nil, false, &os.PathError{
			Op:   "stat",
			Path: path,
			Err:  errors.New("path is a directory"),
		}
	}

	return stat, true, nil
}
