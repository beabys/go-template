package database

import (
	"gitlab.com/beabys/quetzal"
)

type Redis struct {
	*quetzal.Redis
}

// RedisConnect return new Redis
func NewRedis(c *quetzal.RedisConfig) *Redis {
	redis := quetzal.NewRedis().SetConfigs(c)
	return &Redis{redis}
}
