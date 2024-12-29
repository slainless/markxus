package config

type KV map[string]any

func (s KV) Lookup(key string) (any, bool) {
	val, ok := s[key]
	return val, ok
}

func (s KV) String() string {
	return "Global YAML config"
}

func (s KV) GoString() string {
	return "&globalYamlSource{}"
}
