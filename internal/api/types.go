package api

import (
	"context"
	"net"
	"net/http"

	"github.com/beabys/go-template/internal/app/config"
	helloworld "github.com/beabys/go-template/internal/hello_world"
	"github.com/beabys/go-template/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	hwproto "github.com/beabys/go-template/proto/gen/go/hello_world/v1"
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
