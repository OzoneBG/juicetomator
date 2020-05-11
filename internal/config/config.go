package config

import "github.com/urfave/cli/v2"

// JuiceConfig holds all configuration variables.
type JuiceConfig struct {
	HTTPPort string
}

// NewJuiceConfig creates a new configuration instance using the provided context.
func NewJuiceConfig(c *cli.Context) *JuiceConfig {
	return &JuiceConfig{
		HTTPPort: c.String("http-port"),
	}
}
