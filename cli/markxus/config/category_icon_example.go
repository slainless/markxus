package config

import (
	_ "embed"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/slainless/markxus"
)

//go:embed skyrim_icon_map.json
var skyrimIconMapRaw string
var skyrimIconMap []markxus.CategoryIconMap

var exampleIconMapInit = sync.Once{}
var exampleIconError error

func exampleIconMap() []markxus.CategoryIconMap {
	exampleIconMapInit.Do(func() {
		exampleIconError = sonic.UnmarshalString(skyrimIconMapRaw, &skyrimIconMap)
	})
	if exampleIconError != nil {
		panic(exampleIconError)
	}

	return skyrimIconMap
}
