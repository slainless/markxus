package fs

import "regexp"

var regexpStripper = regexp.MustCompile(`[\\/:*?"<>|\x00-\x1F]`)

func Stripper(name string) string {
	return regexpStripper.ReplaceAllString(name, "")
}
