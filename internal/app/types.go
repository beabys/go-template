package app

import (
	"github.com/beabys/go-template/internal/app/ports"
	grpcdapter "github.com/beabys/go-template/internal/infrastructure/adapters/grpc"
	httpadapter "github.com/beabys/go-template/internal/infrastructure/adapters/http"
	"github.com/beabys/go-template/pkg/database"
	"github.com/beabys/go-template/pkg/logger"
)

// App is the Application Struct
type App struct {
	Config      ports.AppConfig
	Logger      logger.Logger
	MysqlClient database.Database
	RedisClient database.Database
	HttpServer  *httpadapter.HttpServer
	GrpcServer  *grpcdapter.GRPCServer
}
