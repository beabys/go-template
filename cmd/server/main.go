package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/beabys/go-template/internal/app"
	"github.com/beabys/go-template/internal/app/config"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
	err := app.Setup(config)
	if err != nil {
		panic(err)
	}

	// Connect to Mysql
	err = app.MysqlClient.Connect()
	if err != nil {
		app.Logger.Fatal("error setting Mysql client", zap.Error(err))
	}

	// // Connect to Redis
	// err = app.RedisClient.Connect()
	// if err != nil {
	// 	app.Logger.Fatal("error setting Redis client", zap.Error(err))
	// }

	// Setup the http Server
	err = app.SetHTTPServer()
	if err != nil {
		app.Logger.Fatal("error setting http server", zap.Error(err))
	}

	// Setup the GRPC Server
	err = app.SetGRPCServer()
	if err != nil {
		app.Logger.Fatal("error setting grpc server", zap.Error(err))
	}

	wg, ctx := errgroup.WithContext(ctx)

	// run servers
	app.HttpServer.Run(ctx, wg)
	app.GrpcServer.Run(ctx, wg)

	err = wg.Wait()
	if err != nil {
		app.Logger.Error("application stopped with error:", err)
	}
	app.Logger.Info("application stopped")
}
