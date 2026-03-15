package grpc

import (
	"context"
	"errors"
	"net"

	"github.com/beabys/go-template/internal/app/config"
	"github.com/beabys/go-template/internal/application/example/command"
	"github.com/beabys/go-template/internal/application/example/handler"
	"github.com/beabys/go-template/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	hwproto "github.com/beabys/go-template/proto/gen/go/hello_world/v1"
)

type GRPCServer struct {
	hwproto.UnimplementedHelloWorldServiceServer
	Server         *grpc.Server
	Listener       net.Listener
	Config         *config.Config
	Logger         logger.Logger
	ExampleService handler.ExampleServiceHandler
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (gs *GRPCServer) SetConfig(c *config.Config) *GRPCServer {
	gs.Config = c
	return gs
}

func (gs *GRPCServer) SetLogger(l logger.Logger) *GRPCServer {
	gs.Logger = l
	return gs
}

func (gs *GRPCServer) SetExampleService(svc handler.ExampleServiceHandler) *GRPCServer {
	gs.ExampleService = svc
	return gs
}

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

func (gs *GRPCServer) GetHelloWorld(ctx context.Context, _ *hwproto.HelloWorldRequest) (*hwproto.HelloWorldResponse, error) {
	resp, err := gs.ExampleService.GetHelloWorld(ctx, &command.GetHelloWorldRequest{})
	if err != nil {
		gs.Logger.Error("error from ExampleService", err)
		return nil, err
	}
	return &hwproto.HelloWorldResponse{
		Hello: resp.Message,
	}, nil
}
