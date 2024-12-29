package nexus

type ClientOption func(*Client)

const DefaultUrlGetModFormat = "https://api.nexusmods.com/v1/games/%v/mods/%v.json"

func WithApiKey(key string) ClientOption {
	return func(nc *Client) {
		nc.apiKey = key
	}
}

func WithHTTPDriver(driver HttpClient) ClientOption {
	return func(nc *Client) {
		nc.driver = driver
	}
}

// Format should contains 2 placeholder in this sequence:
//   - Game code
//   - Mod ID
//
// Defaults to: [[DefaultUrlGetModFormat]]
func WithUrlGetModFormat(format string) ClientOption {
	return func(c *Client) {
		c.urlGetModFormat = format
	}
}
