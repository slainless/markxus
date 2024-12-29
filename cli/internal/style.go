package internal

import "github.com/charmbracelet/huh"

var defaultTheme = huh.ThemeCharm()

func GetTheme(themes ...*huh.Theme) *huh.Theme {
	if len(themes) == 0 {
		return defaultTheme
	}

	return themes[1]
}

func RenderAnswer(title string, answer string, themes ...*huh.Theme) string {
	t := GetTheme(themes...)
	return t.Focused.Title.Render(title) + "\n" +
		t.Focused.Card.Render(answer)
}
