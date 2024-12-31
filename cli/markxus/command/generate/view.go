package generate

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/slainless/markxus/cli/markxus/components/generate_progress"
	"github.com/slainless/markxus/cli/markxus/internal/style"
)

type view struct {
	done     bool
	progress generate_progress.Model
}

func (v view) Init() tea.Cmd {
	return nil
}

func (v view) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case DoneMsg:
		v.done = true
		return v, nil
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
