package colorizer

import (
	"fmt"
	"image/color"
	"sync"
)

var rgbaCache = map[string]*color.RGBA{}
var rgbaMutex = sync.Mutex{}

type OkLCH struct {
	L float64
	C float64
	H float64
	A uint8
}

func (c *OkLCH) key() string {
	return fmt.Sprintf("%v%v%v", c.L, c.C, c.H)
}

func (c *OkLCH) RGBA() *color.RGBA {
	if rgba := rgbaCache[c.key()]; rgba != nil {
		return &color.RGBA{R: rgba.R, G: rgba.G, B: rgba.B, A: c.A}
	}

	rgbaMutex.Lock()
	defer rgbaMutex.Unlock()

	if rgba := rgbaCache[c.key()]; rgba != nil {
		return &color.RGBA{R: rgba.R, G: rgba.G, B: rgba.B, A: c.A}
	}

	rgba := okLCHToRGBA(c)
	rgba.A = 255
	rgbaCache[c.key()] = rgba
	return &color.RGBA{R: rgba.R, G: rgba.G, B: rgba.B, A: c.A}
}
