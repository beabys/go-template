package api

import (
	"context"
	"errors"

	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/go-http-template/internal/utils"
	"gitlab.com/beabys/go-http-template/pkg/logger"
)

// NewHttpServer returns a new pointer of HttpServer
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

// SetConfig is a setter function to set Configs
func (gs *GRPCServer) SetConfig(c *config.Config) *GRPCServer {
	gs.Config = c
	return gs
}

// SetLogger is a setter function to set the Logger
func (gs *GRPCServer) SetLogger(l logger.Logger) *GRPCServer {
	gs.Logger = l
	return gs
}

// SetHelloWorldService is a setter function to set the Logger
func (gs *GRPCServer) SetHelloWorldService(hw helloworld.HelloWorldIntereface) *GRPCServer {
	gs.HelloWorldSvc = hw
	return gs
}

// Run implements Run apoi server function for Htt server
func (gs *GRPCServer) Run(ctx context.Context, cancelFn context.CancelFunc) error {
	go func() {
		if err := gs.Server.Serve(gs.Listener); err != nil {
			err = utils.BindError(errors.New("grpc server stopped with error"), err)
			gs.Logger.Fatal(err)
		}
	}()

	gs.Logger.Info("app started")

	<-ctx.Done()
	cancelFn()
	gs.Logger.Info("shutting down gracefully start")
	gs.Server.GracefulStop()
	return nil
}
