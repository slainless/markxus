package edit

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/slainless/markxus/cli/markxus/command/config/internal"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/fs"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	configType := internal.ConfigType()
	configPath := internal.ConfigPath(configType)

	path := config.ConfigPath(configType)
	stat, exist, err := fs.IsFileExist(path)
	if err != nil {
		return err
	}

	if !exist {
		if !internal.IsOverwrite() {
			return fmt.Errorf("config file not found.\nInit first with *markxus config init -t %s*", configType)
		}

		err := internal.WriteToFile(configPath, config.CreateDefaultConfig())
		if err != nil {
			return err
		}
	} else {
		mode := stat.Mode()
		if mode&os.ModePerm == 0 || (mode&0400 == 0 && mode&0040 == 0) {
			return fmt.Errorf("no permission to read:\n%s", configPath)
		}
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", configPath)
	case "darwin": // macOS
		cmd = exec.Command("open", configPath)
	case "linux":
		cmd = exec.Command("xdg-open", configPath)
	default:
		return fmt.Errorf("unsupported operating system")
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	theme := style.GetTheme()
	fmt.Println(style.Card().Render(
		fmt.Sprintf(
			"Launched file editor\n%s",
			theme.Focused.Description.Render(configPath),
		),
	))
	return nil
}
