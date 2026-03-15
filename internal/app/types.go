package app

import (
	"context"

	"github.com/beabys/go-template/internal/app/config"
	grpcdapter "github.com/beabys/go-template/internal/infrastructure/adapters/grpc"
	httpadapter "github.com/beabys/go-template/internal/infrastructure/adapters/http"
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
	HttpServer  *httpadapter.HttpServer
	GrpcServer  *grpcdapter.GRPCServer
}
