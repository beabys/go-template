package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"gitlab.com/beabys/go-http-template/internal/api"
	"gitlab.com/beabys/go-http-template/internal/app/config"
	"gitlab.com/beabys/go-http-template/internal/app/database"
	"gitlab.com/beabys/go-http-template/internal/app/handler"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/go-http-template/internal/utils"
	"gitlab.com/beabys/go-http-template/pkg/logger"
	"gitlab.com/beabys/quetzal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	hwproto "gitlab.com/beabys/go-http-template/proto/gen/go/hello_world/v1"
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

func (app *App) SetMysqlClient(m *database.Mysql) {
	app.MysqlClient = m
}

func (app *App) SetRedisClient(r *database.Redis) {
	app.RedisClient = r
}

func (app *App) Run(ctx context.Context) error {
	return app.ApiServer.Run(ctx, app.StopFn)
}

func (app *App) Setup(configs config.AppConfig, stopFn context.CancelFunc) error {
	app.StopFn = stopFn
	err := app.SetConfigs(configs)
	if err != nil {
		return err
	}

	config := app.Config.GetConfigs()

	// SetLogger
	loggerConfigs := &logger.DefaultLoggerConfig{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
	}
	logger := logger.NewDefaultLogger(loggerConfigs)
	app.SetLogger(logger)

	// Mysql Client
	mysqlConfig := &quetzal.MysqlConfig{
		Username:        config.DB.Username,
		Password:        config.DB.Password,
		Host:            config.DB.Host,
		Port:            config.DB.Port,
		DBName:          config.DB.DBName,
		LogSQL:          config.DB.LogSQL,
		MaxIdleConns:    config.DB.MaxIdleConns,
		MaxOpenConn:     config.DB.MaxOpenConn,
		ConnMaxLifetime: config.DB.ConnMaxLifetime,
	}
	mysql := database.NewMysql(mysqlConfig)
	app.SetMysqlClient(mysql)

	//Redis
	redisConfig := &quetzal.RedisConfig{
		Host:     config.Redis.Host,
		Password: config.Redis.Password,
		Port:     config.Redis.Port,
		DBNumber: config.Redis.DBNumber,
	}
	redis := database.NewRedis(redisConfig)
	app.SetRedisClient(redis)

	return nil
}

func (a *App) SetHTTPServer() {
	// init service dependencies here
	helloWorldService := helloworld.NewHelloWorld(a.Logger)

	configs := a.Config.GetConfigs()
	server := api.NewHttpServer().
		SetConfig(configs).
		SetLogger(a.Logger).
		// TODO this should be changes according new implementations
		SetHelloWorldService(helloWorldService)

	h := handler.NewMuxHandler(server)

	address := fmt.Sprintf("%s:%v", configs.Http.Host, configs.Http.Port)

	a.Logger.Info("setup http server", address)

	httpServer := &http.Server{
		Addr:              address,
		Handler:           h,
		ReadHeaderTimeout: time.Duration(30 * 1000),
	}
	server.Server = httpServer
	a.ApiServer = server
}

func (a *App) SetGRPCServer() error {
	// init service dependencies here
	helloWorldService := helloworld.NewHelloWorld(a.Logger)

	configs := a.Config.GetConfigs()
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(configs.Grpc.Port))
	if err != nil {
		return utils.BindError(errors.New("failed to get grpc listener"), err)
	}
	server := api.NewGRPCServer().
		SetConfig(configs).
		SetLogger(a.Logger).

		// TODO this should be changes according new implementations
		SetHelloWorldService(helloWorldService)

	recoveryOpt := grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		err = utils.BindError(errors.New("panic recovered"), err)
		a.Logger.Error(err)
		return err
	})
	// get logger and create new entry for grpcLogger``
	logrusLogger := a.Logger.GetLogger().(*logrus.Logger)
	grpcLogger := logrus.NewEntry(logrusLogger)

	rpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(grpcLogger),
			grpc_recovery.StreamServerInterceptor(recoveryOpt),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(grpcLogger),
			grpc_recovery.UnaryServerInterceptor(recoveryOpt),
		)),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	hwproto.RegisterHelloWorldServiceServer(rpcServer, server)
	reflection.Register(rpcServer)

	server.Listener = listener
	server.Server = rpcServer

	a.ApiServer = server
	return nil
}

// Recoverer is a recover function that allow restart the service
// log the error and the stack trace
func (app *App) Recoverer(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			logger := app.GetLogger()
			stackTrace := fmt.Sprintf("%v\n%v", r, string(debug.Stack()))
			message := "Recovering from panic"
			logger.Error(message, fmt.Errorf(stackTrace))
			go app.Recoverer(fn)
		}
	}()
	fn()
}
