package generate_progress

import (
	"math"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/slainless/markxus/nexus"
)

type StartMsg int
type ModDiscoveredMsg *nexus.SchemaMod
type GenerationProgressMsg int
type DoneMsg int
type ErrorMsg error

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case StartMsg:
		if m.status != StatusQueued {
			return m, nil
		}

		m.status = StatusStarted
		return m, nil

	case ModDiscoveredMsg:
		if m.status != StatusStarted {
			return m, nil
		}

		m.mod = msg
		m.status = StatusModDiscovered
		m.progress = 15
		cmd := m.bar.SetPercent(m.progress)
		return m, cmd

	case GenerationProgressMsg:
		if m.status != StatusModDiscovered && m.status != StatusGeneratingMarkdown {
			return m, nil
		}

		m.status = StatusGeneratingMarkdown
		m.progress = increment(m.progress, 0.05)
		cmd := m.bar.SetPercent(float64(m.progress))
		return m, cmd

	case DoneMsg:
		if m.status != StatusModDiscovered && m.status != StatusGeneratingMarkdown {
			return m, nil
		}

		m.status = StatusDone
		m.progress = 100
		cmd := m.bar.SetPercent(100)
		return m, cmd

	case ErrorMsg:
		m.status = StatusError
		m.err = msg
		return m, nil

	case progress.FrameMsg:
		bar, cmd := m.bar.Update(msg)
		if bar, ok := bar.(progress.Model); ok {
			m.bar = bar
		}
		return m, cmd

	}

	return m, nil
}

func increment(current float64, decayFactor float64) float64 {
	progress := current + ((100 - current) * math.Exp(-decayFactor))

	if progress >= 99.8 {
		progress = 99.8
	}

	return progress
}
