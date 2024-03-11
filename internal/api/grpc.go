package api

import (
	"context"

	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/go-http-template/pkg/logger"
	"go.uber.org/zap"
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

// Run implements Run api server function for gRPC server
func (gs *GRPCServer) Run(ctx context.Context, cancelFn context.CancelFunc) error {
	go func() {
		if err := gs.Server.Serve(gs.Listener); err != nil {
			gs.Logger.Fatal("grpc server stopped with error", zap.Error(err))
		}
	}()

	gs.Logger.Info("app started")

	<-ctx.Done()
	cancelFn()
	gs.Logger.Info("shutting down gracefully start")
	gs.Server.GracefulStop()
	return nil
}
