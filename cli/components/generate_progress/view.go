package generate_progress

import (
	"fmt"

	"github.com/slainless/markxus/cli/markxus/components/mod_name"
)

func (m Model) View() string {
	title := mod_name.View(m.GameCode, m.ModId)

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
