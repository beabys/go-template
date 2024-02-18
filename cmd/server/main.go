package main

import (
	"time"

	"gitlab.com/beabys/go-http-template/internal/app"
	"gitlab.com/beabys/go-http-template/internal/app/config"
)

func main() {
	// First we Set a new App
	// New App
	app := app.New()

	// Setup configurations
	config := config.New()
	err := app.Setup(config)
	if err != nil {
		panic(err)
	}

	// Connect to Mysql
	err = app.MysqlClient.Connect()
	if err != nil {
		panic(err)
	}

	// Connect to Redis
	err = app.RedisClient.Connect()
	if err != nil {
		panic(err)
	}

	// start mux server
	// Server already has a Recovery middleware
	go app.Router.Serve(config.App.Host, config.App.Port)

	// Blocking until the shutdown is completed
	<-app.ChanInterrupt
	time.Sleep(time.Second * time.Duration(10))
	app.Logger.Info("Shutdown Completed")
}
