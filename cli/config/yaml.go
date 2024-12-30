package config

import (
	"os"
	"path"

	"github.com/goccy/go-yaml"
)

var (
	ConfigPathGlobal string
	ConfigPathLocal  string
)

var (
	YamlSourceGlobal KV
	YamlSourceLocal  KV
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ConfigPathGlobal = path.Join(home, ".markxus.yml")
	ConfigPathLocal = path.Join(cwd, ".markxus.yml")

	YamlSourceGlobal = NewYamlSource(ConfigPathGlobal)
	YamlSourceLocal = NewYamlSource(ConfigPathLocal)
}

func NewYamlSource(path string) KV {
	kv := KV{}
	config, err := os.ReadFile(ConfigPathGlobal)
	if err != nil {
		return kv
	}

	_ = yaml.Unmarshal(config, &kv)
	return kv
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
