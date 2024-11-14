package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewRedis(c *RedisConfig) *Redis {
	r := &Redis{}
	r.config = c
	return r
}

// Connect Create the connection with Redis
func (r *Redis) Connect() error {
	if r.config.Host == "" {
		r.config.Host = "127.0.0.1"
	}
	if r.config.Port == 0 {
		r.config.Port = 6379
	}
	config := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.config.Host, r.config.Port), // use default Addr
		DB:   r.config.DBNumber,                                  // use default DB
	}
	if r.config.Password != "" {
		config.Password = r.config.Password
	}

	db := redis.NewClient(config)

	_, err := db.Ping().Result()
	if err != nil {
		return err
	}

	r.DB = db
	return nil
}

// Ping verifies a connection to the database is still alive
func (r *Redis) Ping() error {
	return r.DB.Ping().Err()
}

// Close closes the database and prevents new queries from starting.
func (r *Redis) Close() error {
	return r.DB.Close()
}

func (r *Redis) GetDBImpl() any {
	return r
}
