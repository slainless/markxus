package colorizer

var white = &OkLCH{L: 1.0, C: 0.0, H: 0.0, A: 255}
var black = &OkLCH{L: 0.0, C: 0.0, H: 0.0, A: 255}

func ForegroundFrom(color *OkLCH) *OkLCH {
	contrastWithBlack := contrastRatio(color.L, black.L)
	contrastWithWhite := contrastRatio(color.L, white.L)

	if contrastWithBlack >= 4.5 {
		return black
	} else if contrastWithWhite >= 4.5 {
		return white
	}

	// should be nil...
	return white
}

func contrastRatio(L1, L2 float64) float64 {
	if L1 < L2 {
		L1, L2 = L2, L1
	}
	return (L1 + 0.05) / (L2 + 0.05)
}
