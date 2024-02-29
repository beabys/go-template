package api

import (
	"context"
	"net"
	"net/http"

	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/go-http-template/pkg/logger"
	"google.golang.org/grpc"

	hwproto "gitlab.com/beabys/go-http-template/proto/gen/go/hello_world/v1"
)

// HttpServer is a struct of an Http Server
type HttpServer struct {
	Server        *http.Server
	Config        *config.Config
	Logger        logger.Logger
	HelloWorldSvc helloworld.HelloWorldIntereface
}

// HttpServer is a struct of an Http Server
type GRPCServer struct {
	Server        *grpc.Server
	Listener      net.Listener
	Config        *config.Config
	Logger        logger.Logger
	HelloWorldSvc helloworld.HelloWorldIntereface

	hwproto.UnimplementedHelloWorldServiceServer
}

type ApiServer interface {
	Run(context.Context, context.CancelFunc) error
}
