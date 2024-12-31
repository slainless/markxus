package internal

import cfg "github.com/slainless/markxus/cli/markxus/config"

func ConfigType() cfg.ConfigType {
	return cfg.Config.Common.ConfigType.Selected()
}

func ConfigPath(t cfg.ConfigType) string {
	return cfg.ConfigPath(t)
}

func IsOverwrite() bool {
	return cfg.Config.Common.Overwrite
}
