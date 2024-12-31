package generate_progress

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/slainless/markxus/cli/markxus/colorizer"
)

func (m Model) View() string {
	gameBg, gameFg := colorizer.GenerateBackground(m.GameCode)
	modBg, modFg := colorizer.GenerateBackground(m.ModId)

	titleStyle := lipgloss.NewStyle().Padding(0, 1)
	title := fmt.Sprint(
		titleStyle.Foreground(lipgloss.Color(gameFg)).Background(lipgloss.Color(gameBg)).Render(m.GameCode),
		titleStyle.Foreground(lipgloss.Color(modFg)).Background(lipgloss.Color(modBg)).Render(m.ModId),
	)

	var modName string
	if m.mod == nil {
		modName = m.Theme.Focused.TextInput.Placeholder.Render("(Not fetched yet)")
	} else {
		modName = m.mod.Name
	}

	var bar string
	switch m.status {
	case StatusQueued:
		bar = m.Theme.Focused.SelectSelector.Render("Queued")
	case StatusGeneratingMarkdown, StatusModDiscovered, StatusStarted:
		bar = m.bar.View()
	case StatusDone:
		bar = m.Theme.Focused.SelectedOption.Render("Done")
	case StatusError:
		bar = m.Theme.Focused.ErrorMessage.Render(m.err.Error())
	}

	return fmt.Sprintf("%s\n%s\n%s",
		title,
		modName,
		bar,
	)
}
