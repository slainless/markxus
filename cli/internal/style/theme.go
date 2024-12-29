package style

import "github.com/charmbracelet/huh"

var defaultTheme = huh.ThemeCharm()

func GetTheme(themes ...*huh.Theme) *huh.Theme {
	if len(themes) == 0 {
		return defaultTheme
	}

	return themes[1]
}
