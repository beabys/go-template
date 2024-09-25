package app

import (
	"context"

	"gitlab.com/beabys/go-template/internal/api"
	"gitlab.com/beabys/go-template/internal/app/config"
	"gitlab.com/beabys/go-template/pkg/database"
	"gitlab.com/beabys/go-template/pkg/logger"
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
	MysqlClient database.Database
	// RedisClient *database.Redis
	HttpServer api.ApiServer
	GrpcServer api.ApiServer
}
