package generate

import (
	"fmt"
	"os"

	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/fs"
)

func checkMarkdown(path string) error {
	stat, exist, err := fs.IsFileExist(path)
	if err != nil {
		return err
	}

	if !exist {
		return nil
	} else {
		if !config.Config.Common.Overwrite {
			return fmt.Errorf("markdown already exist in path %s.\nRun the command with overwrite flag", path)
		}

		mode := stat.Mode()
		if mode&os.ModePerm == 0 || (mode&0200 == 0 && mode&0020 == 0) {
			return fmt.Errorf("no permission to write to:\n%s", path)
		}
	}

	return nil
}
