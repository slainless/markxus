package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var GlobalConfigDir string
var LocalConfigDir string

var GlobalYamlSource KV
var LocalYamlSource KV

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	GlobalConfigDir = path.Join(home, ".markxus.yml")
	LocalConfigDir = path.Join(cwd, ".markxus.yml")

	GlobalYamlSource = NewYamlSource(GlobalConfigDir)
	LocalYamlSource = NewYamlSource(LocalConfigDir)
}

func NewYamlSource(path string) KV {
	kv := KV{}
	config, err := os.ReadFile(GlobalConfigDir)
	if err != nil {
		return kv
	}

	_ = yaml.Unmarshal(config, &kv)
	return kv
}
