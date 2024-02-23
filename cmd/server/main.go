package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/beabys/go-http-template/internal/app"
	"gitlab.com/beabys/go-http-template/internal/app/config"
)

func main() {

	// First we Set a context and a stopFn
	ctx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stopFn()

	// New App
	app := app.New()

	// Setup configurations
	config := config.New()

	//setup the app
	err := app.Setup(config, stopFn)
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
	err = app.Run(ctx)
	if err != nil {
		app.Logger.Error("application stopped with error:", err.Error())
	}
	app.Logger.Info("application stopped")
}
