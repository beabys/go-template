package app

import (
	"context"

	"gitlab.com/beabys/go-http-template/internal/api"
	"gitlab.com/beabys/go-http-template/internal/app/config"
	"gitlab.com/beabys/go-http-template/internal/app/database"
	"gitlab.com/beabys/go-http-template/pkg/logger"
)

type Application interface {
	SetLogger(logger.Logger)
	Run(context.Context) error
	Setup(config.AppConfig, context.CancelFunc) error
	Recoverer(func())
}

// App is the Application Struct
type App struct {
	Config      config.AppConfig
	Logger      logger.Logger
	StopFn      context.CancelFunc
	MysqlClient *database.Mysql
	RedisClient *database.Redis
	ApiServer   api.ApiServer
}
