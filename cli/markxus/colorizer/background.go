package colorizer

import (
	"crypto/sha256"
	"sync"
)

var cache = map[string]*OkLCH{}
var cacheMutex = sync.Mutex{}

const mask = 0xFFFF

func BackgroundFrom(text string) *OkLCH {
	if c := cache[text]; c != nil {
		return c
	}

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if c := cache[text]; c != nil {
		return c
	}

	hash := sha256.New()
	hash.Write([]byte(text))
	hashBytes := hash.Sum(nil)

	hashInt := binaryToUint64(hashBytes[:8])

	color := &OkLCH{
		L: normalize(hashInt&mask, 1),         // Lightness: [0.0, 1.0]
		C: normalize(hashInt>>16&mask, 0.8),   // Chroma: [0.0, 0.8]
		H: normalize(hashInt>>32&mask, 360.0), // Hue: [0.0, 360.0]
		A: 255,
	}
	cache[text] = color
	return color
}

func normalize(hashValue uint64, max float64) float64 {
	return (max * float64(hashValue)) / mask
}

func binaryToUint64(bytes []byte) uint64 {
	var result uint64
	for i := 0; i < 8; i++ {
		result |= uint64(bytes[i]) << (8 * (7 - i))
	}
	return result
}
