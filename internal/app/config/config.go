package config

import (
	config "gitlab.com/beabys/ayotl"
)

var CONFIG_FILE = "CONFIG_FILE"

// New return  a New Config
func New() *Config {
	return &Config{}
}

func (c *Config) GetConfigs() *Config {
	return c
}

// LoadConfig is a function to load the configuration, stored on the config files
func (c *Config) LoadConfigs() error {
	bc := config.New().SetConfigImpl(c).WithEnv(CONFIG_FILE)
	if err := bc.LoadConfigs(c, bc.MustString(CONFIG_FILE, "")); err != nil {
		return err
	}
	return nil
}

func (c *Config) SetDefaults() config.ConfigMap {
	defaults := make(config.ConfigMap)
	return defaults
}
