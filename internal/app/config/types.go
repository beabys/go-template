package config

type AppConfig interface {
	LoadConfigs() error
	GetConfigs() *Config
}

// Config is a struct define configuration for the app
type Config struct {
	Stage string   `mapstructure:"stage"`
	Http  Http     `mapstructure:"http"`
	Grpc  Grpc     `mapstructure:"grpc"`
	DB    Database `mapstructure:"db"`
	Redis Redis    `mapstructure:"redis"`
}

// Grpc is a struct to define configurations for the gRPC Server
type Grpc struct {
	Port int `mapstructure:"port"`
}

// Http is a struct to define configurations for the http Server
type Http struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// MysqlConfigurations is a struct to define configurations of the db connection
type Database struct {
	Username        string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DBName          string `mapstructure:"name"`
	LogSQL          bool   `mapstructure:"log_sql"`
	MaxIdleConns    int    `mapstructure:"max_idle_connections"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	ConnMaxLifetime int    `mapstructure:"connection_max_lifetime_seconds"`
}

// Redis is a struct to define configurations for Redis
type Redis struct {
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBNumber int    `mapstructure:"database"`
}
