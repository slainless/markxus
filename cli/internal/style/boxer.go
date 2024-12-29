package style

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func Card(theme ...*huh.Theme) lipgloss.Style {
	t := GetTheme(theme...)
	return lipgloss.NewStyle().
		Padding(1, 2).
		Width(40).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.FieldSeparator.GetForeground())
}
