package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/beabys/go-http-template/internal/api"
	"gitlab.com/beabys/go-http-template/internal/app/config"
	"gitlab.com/beabys/go-http-template/internal/app/database"
	"gitlab.com/beabys/go-http-template/internal/app/handler"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/quetzal"
	"go.uber.org/zap"
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

func (app *App) SetLogger(logger quetzal.Logger) {
	app.Logger = logger
}

func (app *App) GetLogger() quetzal.Logger {
	return app.Logger
}

func (app *App) SetMysqlClient(m *database.Mysql) {
	app.MysqlClient = m
}

func (app *App) SetRedisClient(r *database.Redis) {
	app.RedisClient = r
}

func (app *App) Run(ctx context.Context) error {
	var err error = nil
	httpServer := app.initHTTPServer(ctx)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			newError := errors.Join(errors.New("http server stopped with error"), err)
			app.Logger.Fatal(newError)
		}
	}()

	app.Logger.Info("app started")

	<-ctx.Done()
	app.StopFn()
	app.Logger.Info("shutting down gracefully start")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = httpServer.Shutdown(ctxTimeout); err != nil {
		app.Logger.Error("error shutting server down", err)
	}
	return err
}

func (app *App) Setup(configs config.AppConfig, stopFn context.CancelFunc) error {
	app.StopFn = stopFn
	err := app.SetConfigs(configs)
	if err != nil {
		return err
	}

	config := app.Config.GetConfigs()

	// SetLogger
	loggerConfigs := &quetzal.DefaultLoggerConfig{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
	}
	logger := quetzal.NewDefaultLogger(loggerConfigs)
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

func (a *App) initHTTPServer(ctx context.Context) *http.Server {
	// init service dependencies here
	helloWorldService := helloworld.NewHelloWorld(a.Logger)

	server := api.NewHttpServer().
		SetConfig(a.Config.GetConfigs()).
		SetLogger(a.Logger).
		SetHelloWorldService(helloWorldService)

	h := handler.NewMuxHandler(ctx, server)

	a.Logger.Info("setup http server", zap.String("port", fmt.Sprintf("%v", 8080)))

	return &http.Server{
		Addr:              fmt.Sprintf("%s:%v", "", 8080),
		Handler:           h,
		ReadHeaderTimeout: time.Duration(30 * 1000),
	}
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
