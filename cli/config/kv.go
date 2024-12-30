package config

import "fmt"

type KV map[string]any

func (s KV) Lookup(key string) (any, bool) {
	val, ok := s[key]
	if val == nil || val == "" {
		return nil, false
	}

	return val, ok
}

func (s KV) String() string {
	return fmt.Sprint(map[string]any(s))
}

func (s KV) GoString() string {
	return fmt.Sprint(map[string]any(s))
}
