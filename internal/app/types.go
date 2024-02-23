package app

import (
	"context"

	"gitlab.com/beabys/go-http-template/internal/app/config"
	"gitlab.com/beabys/go-http-template/internal/app/database"
	"gitlab.com/beabys/quetzal"
)

type Application interface {
	SetLogger(quetzal.Logger)
	Run(context.Context) error
	Setup(config.AppConfig, context.CancelFunc) error
	Recoverer(func())
}

// App is the Application Struct
type App struct {
	Config      config.AppConfig
	Logger      quetzal.Logger
	StopFn      context.CancelFunc
	MysqlClient *database.Mysql
	RedisClient *database.Redis
}
