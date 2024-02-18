package app

import (
	"fmt"
	"os"
	"runtime/debug"

	m "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"gitlab.com/beabys/go-http-template/internal/app/config"
	"gitlab.com/beabys/go-http-template/internal/app/database"
	"gitlab.com/beabys/go-http-template/internal/app/utils"
	"gitlab.com/beabys/go-http-template/pkg/router"
	"gitlab.com/beabys/quetzal"
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

func (app *App) SetMuxRouter(m router.Router) {
	app.Router = m
}

func (app *App) SetMysqlClient(m *database.Mysql) {
	app.MysqlClient = m
}

func (app *App) SetRedisClient(r *database.Redis) {
	app.RedisClient = r
}

func (app *App) SetChanInterrupt(f chan interface{}) {
	app.ChanInterrupt = f
}

func (app *App) Run() {
	app.Logger.Info("let's Run")
}

func (app *App) Setup(configs config.AppConfig) error {
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

	// Set Mux Router
	muxRouter := router.
		NewDefaultRouter().
		SetLogger(app.Logger)
	muxRouter.SetDefaultMiddlewares()
	muxRouter.Mux.Use(m.Logger)
	app.SetMuxRouter(muxRouter)

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

	app.SetChanInterrupt(utils.InterruptCh(logger, "Main"))
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
