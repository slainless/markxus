package nexus

type ClientOptions struct {
	Driver HttpClient
	ApiKey string

	UrlGetModFormat   string
	UrlGetFilesFormat string
}

type ClientOption func(*ClientOptions)

const DefaultUrlGetModFormat = "https://api.nexusmods.com/v1/games/%v/mods/%v.json"
const DefaultUrlGetFilesFormat = "https://api.nexusmods.com/v1/games/%v/mods/%v/files.json"

func WithApiKey(key string) ClientOption {
	return func(nc *ClientOptions) {
		nc.ApiKey = key
	}
}

func WithHTTPDriver(driver HttpClient) ClientOption {
	return func(nc *ClientOptions) {
		nc.Driver = driver
	}
}

// Format should contains 2 placeholder in this sequence:
//   - Game code
//   - Mod ID
//
// Defaults to: [[DefaultUrlGetModFormat]]
func WithUrlGetModFormat(format string) ClientOption {
	return func(c *ClientOptions) {
		c.UrlGetModFormat = format
	}
}

func WithUrlGetFilesFormat(format string) ClientOption {
	return func(c *ClientOptions) {
		c.UrlGetFilesFormat = format
	}
}
