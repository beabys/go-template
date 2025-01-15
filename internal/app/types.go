package app

import (
	"context"

	"github.com/beabys/go-template/internal/api"
	"github.com/beabys/go-template/internal/app/config"
	"github.com/beabys/go-template/pkg/database"
	"github.com/beabys/go-template/pkg/logger"
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
	RedisClient database.Database
	HttpServer  api.ApiServer
	GrpcServer  api.ApiServer
}
