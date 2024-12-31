package mod_name

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/colorizer"
)

func View(gameCode string, modId string) string {
	gameBg, gameFg := colorizer.GenerateBackground(gameCode)
	modBg, modFg := colorizer.GenerateBackground(modId)

	titleStyle := lipgloss.NewStyle().Padding(0, 1)
	return fmt.Sprint(
		titleStyle.Foreground(lipgloss.Color(gameFg)).Background(lipgloss.Color(gameBg)).Render(gameCode),
		titleStyle.Foreground(lipgloss.Color(modFg)).Background(lipgloss.Color(modBg)).Render(modId),
	)
}
