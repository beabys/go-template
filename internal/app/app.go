package app

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/beabys/go-template/internal/app/config"
	"github.com/beabys/go-template/internal/application/example/handler"
	grpcdapter "github.com/beabys/go-template/internal/infrastructure/adapters/grpc"
	httpadapter "github.com/beabys/go-template/internal/infrastructure/adapters/http"
	"github.com/beabys/go-template/internal/infrastructure/persistence/repository"
	"github.com/beabys/go-template/internal/utils"
	"github.com/beabys/go-template/pkg/database"
	"github.com/beabys/go-template/pkg/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	hwproto "github.com/beabys/go-template/proto/gen/go/hello_world/v1"
)

func New() *App {
	return &App{}
}

func (app *App) SetConfigs(AppConfig config.AppConfig) error {
	app.Config = AppConfig
	err := AppConfig.LoadConfigs()
	if err != nil {
		return err
	}
	return nil
}

func (app *App) SetLogger(logger logger.Logger) {
	app.Logger = logger
}

func (app *App) GetLogger() logger.Logger {
	return app.Logger
}

func (app *App) SetMysqlClient(m database.Database) {
	app.MysqlClient = m
}

func (app *App) Setup(configs config.AppConfig) error {
	err := app.SetConfigs(configs)
	if err != nil {
		return err
	}

	config := app.Config.GetConfigs()

	logger, err := logger.NewZapLogger([]string{}, []string{}, zap.DebugLevel)
	if err != nil {
		return err
	}

	app.SetLogger(logger)

	mysqlConfig := &database.MysqlConfig{
		Username:        config.DB.Username,
		Password:        config.DB.Password,
		Host:            config.DB.Host,
		Port:            config.DB.Port,
		DBName:          config.DB.DBName,
		LogSQL:          config.DB.LogSQL,
		MaxIdleConns:    config.DB.MaxIdleConns,
		MaxOpenConn:     config.DB.MaxOpenConn,
		ConnMaxLifetime: time.Duration(config.DB.ConnMaxLifetime) * time.Second,
	}
	mysql := database.New().SetConfigs(mysqlConfig)

	app.SetMysqlClient(mysql)

	return nil
}

func (a *App) SetHTTPServer() error {
	exampleRepository := repository.NewHelloWorldRepository(a.Logger, a.MysqlClient)
	exampleService := handler.NewExampleService(a.Logger, exampleRepository)

	configs := a.Config.GetConfigs()
	server := httpadapter.NewHttpServer().
		SetConfig(configs).
		SetLogger(a.Logger).
		SetExampleService(exampleService)

	h, err := httpadapter.NewMuxHandler(server)
	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%v", configs.Http.Host, configs.Http.Port)

	a.Logger.Info("setup http server ", logger.LogField{Key: "address", Value: address})

	server.Server = &http.Server{
		Addr:              address,
		Handler:           h,
		ReadHeaderTimeout: time.Duration(time.Second / 2),
	}

	a.HttpServer = server

	return nil
}

func (a *App) SetGRPCServer() error {
	exampleRepository := repository.NewHelloWorldRepository(a.Logger, a.MysqlClient)
	exampleService := handler.NewExampleService(a.Logger, exampleRepository)

	configs := a.Config.GetConfigs()
	address := ":" + strconv.Itoa(configs.Grpc.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return utils.BindError(errors.New("failed to get grpc listener"), err)
	}
	server := grpcdapter.NewGRPCServer().
		SetConfig(configs).
		SetLogger(a.Logger).
		SetExampleService(exampleService)

	recoveryOpt := grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		a.Logger.Error("panic recovered", err)
		return err
	})
	logs := a.Logger.GetLogger()
	opts := []grpc_logging.Option{
		grpc_logging.WithLogOnEvents(grpc_logging.StartCall, grpc_logging.FinishCall),
	}

	rpcServer := grpclib.NewServer(
		grpclib.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logging.StreamServerInterceptor(logger.InterceptorLogger(logs), opts...),
			grpc_recovery.StreamServerInterceptor(recoveryOpt),
		)),
		grpclib.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logging.UnaryServerInterceptor(logger.InterceptorLogger(logs), opts...),
			grpc_recovery.UnaryServerInterceptor(recoveryOpt),
		)),
		grpclib.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	hwproto.RegisterHelloWorldServiceServer(rpcServer, server)
	reflection.Register(rpcServer)

	a.Logger.Info("setup gRPC server ", logger.LogField{Key: "address", Value: address})
	server.Listener = listener
	server.Server = rpcServer

	a.GrpcServer = server
	return nil
}

func (app *App) Recoverer(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			log := app.GetLogger()
			stackTrace := fmt.Sprintf("%v\n%v", r, string(debug.Stack()))
			log.Warn("Recovering from panic", logger.LogField{Key: "stackTrace", Value: stackTrace})
			go app.Recoverer(fn)
		}
	}()
	fn()
}
