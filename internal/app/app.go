package app

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/beabys/go-template/internal/api"
	"github.com/beabys/go-template/internal/app/config"
	helloworld "github.com/beabys/go-template/internal/hello_world"
	"github.com/beabys/go-template/internal/hello_world/repository"
	"github.com/beabys/go-template/internal/utils"
	"github.com/beabys/go-template/pkg/database"
	"github.com/beabys/go-template/pkg/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	hwproto "github.com/beabys/go-template/proto/gen/go/hello_world/v1"
)

// New returns a new App struct
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

// func (app *App) SetRedisClient(r *database.Redis) {
// 	app.RedisClient = r
// }

func (app *App) Setup(configs config.AppConfig) error {
	err := app.SetConfigs(configs)
	if err != nil {
		return err
	}

	config := app.Config.GetConfigs()

	// SetLogger
	loggerConfigs := &logger.DefaultLoggerConfig{}
	logger, err := logger.NewDefaultLogger(loggerConfigs)
	if err != nil {
		return err
	}

	app.SetLogger(logger)

	// Mysql Client
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

	// //Redis
	// redisConfig := &database.RedisConfig{
	// 	Host:     config.Redis.Host,
	// 	Password: config.Redis.Password,
	// 	Port:     config.Redis.Port,
	// 	DBNumber: config.Redis.DBNumber,
	// }
	// redis := database.NewRedis(redisConfig)
	// app.SetRedisClient(redis)

	return nil
}

func (a *App) SetHTTPServer() error {

	// init service dependencies here
	helloWorldRepository := repository.NewHelloRepository(a.Logger, a.MysqlClient)
	helloWorldService := helloworld.NewHelloWorld(a.Logger, helloWorldRepository)

	configs := a.Config.GetConfigs()
	server := api.NewHttpServer().
		SetConfig(configs).
		SetLogger(a.Logger).
		// TODO this should be changes according new implementations
		SetHelloWorldService(helloWorldService)

	h, err := api.NewMuxHandler(server)
	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%v", configs.Http.Host, configs.Http.Port)

	a.Logger.Info("setup http server ", zap.String("address", address))

	server.Server = &http.Server{
		Addr:    address,
		Handler: h,
		// by default set ReadHeaderTimeout to 0.5 secs.
		ReadHeaderTimeout: time.Duration(time.Second / 2),
	}

	a.HttpServer = server

	return nil
}

func (a *App) SetGRPCServer() error {
	// init service dependencies here
	helloWorldRepository := repository.NewHelloRepository(a.Logger, a.MysqlClient)
	helloWorldService := helloworld.NewHelloWorld(a.Logger, helloWorldRepository)

	configs := a.Config.GetConfigs()
	address := ":" + strconv.Itoa(configs.Grpc.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return utils.BindError(errors.New("failed to get grpc listener"), err)
	}
	server := api.NewGRPCServer().
		SetConfig(configs).
		SetLogger(a.Logger).

		// TODO this should be changes according new implementations
		SetHelloWorldService(helloWorldService)

	recoveryOpt := grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		a.Logger.Error("panic recovered", err)
		return err
	})
	// get logger and create new entry for grpcLogger
	logger := a.Logger.GetLogger().(*zap.Logger)

	rpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(recoveryOpt),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(recoveryOpt),
		)),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	hwproto.RegisterHelloWorldServiceServer(rpcServer, server)
	reflection.Register(rpcServer)

	a.Logger.Info("setup gRPC server ", zap.String("address", address))
	server.Listener = listener
	server.Server = rpcServer

	a.GrpcServer = server
	return nil
}

// Recoverer is a recover function that allow restart the service
// log the error and the stack trace
func (app *App) Recoverer(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			logger := app.GetLogger()
			stackTrace := fmt.Sprintf("%v\n%v", r, string(debug.Stack()))
			logger.Warn("Recovering from panic", zap.String("stackTrace", stackTrace))
			go app.Recoverer(fn)
		}
	}()
	fn()
}
