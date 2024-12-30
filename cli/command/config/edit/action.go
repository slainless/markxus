package edit

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/slainless/markxus/cli/markxus/command/config/internal"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	configType := internal.ConfigType()
	configPath := internal.ConfigPath(configType)

	exist, err := isConfigExist(configType)
	if err != nil {
		return err
	}

	if !exist {
		if !internal.IsOverwrite() {
			return fmt.Errorf("config file not found.\nInit first with *markxus config init -t %s*", configType)
		}

		err := internal.WriteToFile(configPath, internal.CreateDefaultConfig())
		if err != nil {
			return err
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
