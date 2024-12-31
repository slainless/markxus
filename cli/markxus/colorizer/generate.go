package colorizer

import (
	"fmt"
	"image/color"
)

func GenerateBackground(text string) (string, string) {
	background := BackgroundFrom(text)
	return hex(background.RGBA()), hex(ForegroundFrom(background).RGBA())
}

func hex(c *color.RGBA) string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}
