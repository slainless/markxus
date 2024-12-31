package generate

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/kennygrant/sanitize"
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/config"
)

func writeMarkdown(generated *markxus.Generated) error {
	outputPath := createOutputPath(generated.Mod.Name)
	fmt.Println(generated.Content)
	time.Sleep(time.Second * 10)
	return os.WriteFile(outputPath, []byte(generated.Header+generated.Content), 0666)
}

func createOutputPath(modName string) string {
	return path.Join(config.Config.Generation.OutputDir, sanitize.Path(modName)+".md")
}
