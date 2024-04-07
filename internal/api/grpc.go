package api

import (
	"context"
	"errors"

	"gitlab.com/beabys/go-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-template/internal/hello_world"
	"gitlab.com/beabys/go-template/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
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
func (gs *GRPCServer) Run(ctx context.Context, wg *errgroup.Group) {
	wg.Go(func() error {
		gs.Logger.Info("grpc server started")
		if err := gs.Server.Serve(gs.Listener); err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			gs.Logger.Error("grpc server failed to serve", err)
			return err
		}
		return nil
	})

	wg.Go(func() error {
		<-ctx.Done()
		gs.Logger.Info("shutting down gracefully grpc server")
		gs.Server.GracefulStop()
		return nil
	})
}
