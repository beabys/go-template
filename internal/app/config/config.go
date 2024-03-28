package config

import (
	config "gitlab.com/beabys/ayotl"
)

// New return  a New Config
func New() *Config {
	return &Config{}
}

func (c *Config) GetConfigs() *Config {
	return c
}

// LoadConfig is a function to load the configuration, stored on the config files
func (c *Config) LoadConfigs() error {
	baseConfig := config.New().
		SetConfigImpl(c).
		LoadEnv()
	configFile := baseConfig.MustString("CONFIG_FILE", "/etc/config/config.yaml")
	err := baseConfig.LoadConfigs(c, configFile)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetDefaults() config.ConfigMap {
	defaults := make(config.ConfigMap)
	return defaults
}
