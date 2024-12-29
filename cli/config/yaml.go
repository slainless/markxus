package config

import (
	"os"
	"path"

	"github.com/goccy/go-yaml"
)

var GlobalConfigPath string
var LocalConfigPath string

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

	GlobalConfigPath = path.Join(home, ".markxus.yml")
	LocalConfigPath = path.Join(cwd, ".markxus.yml")

	GlobalYamlSource = NewYamlSource(GlobalConfigPath)
	LocalYamlSource = NewYamlSource(LocalConfigPath)
}

func NewYamlSource(path string) KV {
	kv := KV{}
	config, err := os.ReadFile(GlobalConfigPath)
	if err != nil {
		return kv
	}

	_ = yaml.Unmarshal(config, &kv)
	return kv
}
