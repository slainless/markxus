package generate_progress

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/huh"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/slainless/markxus/nexus"
)

type Model struct {
	GameCode string
	ModId    string
	Theme    *huh.Theme

	progress float64
	status   Status
	err      error
	mod      *nexus.SchemaMod

	bar progress.Model
}

type Status int

const (
	StatusQueued Status = iota
	StatusStarted
	StatusModDiscovered
	StatusGeneratingMarkdown
	StatusDone
	StatusError
)

func New(gameCode string, modId string, theme ...*huh.Theme) Model {
	t := style.GetTheme(theme...)
	bar := progress.New()

	return Model{
		GameCode: gameCode,
		ModId:    modId,
		Theme:    t,

		progress: 0,
		status:   StatusQueued,
		err:      nil,
		mod:      nil,

		bar: bar,
	}
}

func (m Model) SetStatus(status Status) Model {
	m.status = status
	return m
}
