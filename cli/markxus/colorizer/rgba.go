package colorizer

import (
	"image/color"
	"math"
)

const (
	α = 0.211455
	β = 0.119192
	γ = 0.950304
)

func okLCHToRGBA(oklch *OkLCH) *color.RGBA {
	if oklch.L == 0.0 && oklch.C == 0.0 {
		return &color.RGBA{R: 0, G: 0, B: 0, A: oklch.A}
	}

	if oklch.L == 1.0 && oklch.C == 0.0 {
		return &color.RGBA{R: 255, G: 255, B: 255, A: oklch.A}
	}

	l, a, b := okLCHToOkLAB(oklch)
	r, g, b := okLABToRGB(l, a, b)

	return &color.RGBA{
		R: uint8(math.Round(correctGamma(r) * 255)),
		G: uint8(math.Round(correctGamma(g) * 255)),
		B: uint8(math.Round(correctGamma(b) * 255)),
	}
}

func okLCHToOkLAB(oklch *OkLCH) (float64, float64, float64) {
	return oklch.L,
		oklch.C * math.Cos(oklch.H*math.Pi/180.0),
		oklch.C * math.Sin(oklch.H*math.Pi/180.0)
}

func okLABToRGB(l, a, b float64) (float64, float64, float64) {
	return (l + α*a) / γ, (l + β*b) / γ, (l - α*a - β*b) / γ
}

func correctGamma(c float64) float64 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return (1.055*math.Pow(c, 1.0/2.4) - 0.055)
}
