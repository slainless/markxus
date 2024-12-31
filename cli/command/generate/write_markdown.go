package generate

import (
	"os"
	"path"
	"strings"

	"github.com/kennygrant/sanitize"
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/config"
)

func writeMarkdown(generated *markxus.Generated) error {
	outputPath := createOutputPath(generated.Mod.Name)
	return os.WriteFile(
		outputPath,
		[]byte(strings.Trim(generated.Header+generated.Content, "\n ")),
		0666,
	)
}

func createOutputPath(modName string) string {
	return path.Join(config.Config.Generation.OutputDir, sanitize.Path(modName)+".md")
}
