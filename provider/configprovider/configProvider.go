package configprovider

import (
	"local-cache/provider"
	"os"
)

func NewConfigProvider() provider.ConfigProvider {
	return &Config{}
}

func (c *Config) GetServerPort() string {
	if c == nil {
		return ""
	}
	return c.Port
}

func (c *Config) GetString(key string) string {
	return os.Getenv(key)
}
