package config

import "github.com/slainless/markxus"

func ConfigCategoryIconMap(gameCode string) []markxus.CategoryIconMap {
	if v := lookup(YamlSourceLocal, gameCode); v != nil {
		return v
	} else if v := lookup(YamlSourceGlobal, gameCode); v != nil {
		return v
	}

	if gameCode == "skyrimspecialedition" {
		return exampleIconMap()
	}

	return []markxus.CategoryIconMap{}
}

func lookup(source KV, gameCode string) []markxus.CategoryIconMap {
	iconKey := YamlKeyCategoryIconMap

	if v, _ := source.Lookup(iconKey); v != nil {
		if v, ok := v.(map[string][]markxus.CategoryIconMap); ok {
			if v := v[gameCode]; v != nil {
				return v
			}
		}
	}

	return nil
}
