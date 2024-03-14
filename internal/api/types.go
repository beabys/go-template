package api

import (
	"context"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/go-http-template/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	hwproto "gitlab.com/beabys/go-http-template/proto/gen/go/hello_world/v1"
)

type PrometheusMetrics struct {
	latency *prometheus.SummaryVec
}

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
	Run(context.Context, *errgroup.Group)
}
