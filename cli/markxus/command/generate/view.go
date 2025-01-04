package generate

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/slainless/markxus/cli/markxus/components/generate_progress"
	"github.com/slainless/markxus/cli/markxus/internal/style"
)

type view struct {
	done     bool
	progress generate_progress.Model
	quitErr  error
}

func (v view) Init() tea.Cmd {
	return nil
}

func (v view) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case DoneMsg:
		v.done = true
		return v, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			v.quitErr = errors.New("Cancelled by CTRL+C")
			return v.Update(DoneMsg(0))
		}
	}

	progress, cmd := v.progress.Update(msg)
	v.progress = progress.(generate_progress.Model)
	return v, cmd
}

func (v view) View() string {
	if v.done {
		return ""
	}

	return style.Card().Render(v.progress.View())
}

type DoneMsg int
