package fs

import (
	"errors"
	"os"
)

func IsFileExist(path string) (os.FileInfo, bool, error) {
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
