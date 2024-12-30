package config

import (
	"os"
	"path"

	"github.com/goccy/go-yaml"
)

var ConfigPathGlobal = (func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(home, ".markxus.yml")
})()

var ConfigPathLocal = (func() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path.Join(cwd, ".markxus.yml")
})()

var YamlSourceGlobal, YamlSourceGlobalError = NewYamlSource(ConfigPathGlobal)
var YamlSourceLocal, YamlSourceLocalError = NewYamlSource(ConfigPathGlobal)

func ResolveFromYaml(configType ConfigType, yamlKey string) any {
	var kv KV
	if configType == ConfigTypeGlobal {
		kv = YamlSourceGlobal
	} else if configType == ConfigTypeLocal {
		kv = YamlSourceLocal
	}

	return kv[yamlKey]
}

func NewYamlSource(path string) (KV, error) {
	kv := KV{}
	config, err := os.ReadFile(ConfigPathGlobal)
	if err != nil {
		return kv, err
	}

	err = yaml.Unmarshal(config, &kv)
	return kv, err
}

func ConfigPath(configType ConfigType) string {
	switch configType {
	case ConfigTypeGlobal:
		return ConfigPathGlobal
	case ConfigTypeLocal:
		return ConfigPathLocal
	}

	return ""
}
