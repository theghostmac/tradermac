package alphavantage

import "time"

const defaultTimeout = 2 * time.Second

// Configuration is the config specification for the AlphaVantage APIClient
type Configuration struct {
	URL     string
	APIKeys map[string]string
	Timeout time.Duration
}

// NewConfiguration instantiates a Configuration from the API URL.
func NewConfiguration(URL string) *Configuration {
	return &Configuration{
		URL:     URL,
		APIKeys: make(map[string]string),
		Timeout: defaultTimeout,
	}
}

func (c *Configuration) CustomTimeout(timeout time.Duration) *Configuration {
	c.Timeout = timeout
	return c
}

// AddKey adds a key, value in the APIClient Configuration.
func (c *Configuration) AddKey(key string, value string) *Configuration {
	c.APIKeys[key] = value
	return c
}
