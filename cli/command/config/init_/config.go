package init_

import "github.com/slainless/markxus/cli/markxus/internal"

var initConfig = struct {
	force      bool
	configType *internal.EnumValue
}{
	configType: &internal.EnumValue{
		Enum:    []string{"global", "local"},
		Default: "global",
	},
}
