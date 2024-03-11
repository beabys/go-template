package api

import (
	"context"
	"net/http"
	"time"

	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/go-http-template/pkg/logger"
	"go.uber.org/zap"
)

// NewHttpServer returns a new pointer of HttpServer
func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

// SetConfig is a setter function to set Configs
func (hs *HttpServer) SetConfig(c *config.Config) *HttpServer {
	hs.Config = c
	return hs
}

// SetLogger is a setter function to set the Logger
func (hs *HttpServer) SetLogger(l logger.Logger) *HttpServer {
	hs.Logger = l
	return hs
}

// SetHelloWorldService is a setter function to set the Logger
func (hs *HttpServer) SetHelloWorldService(hw helloworld.HelloWorldIntereface) *HttpServer {
	hs.HelloWorldSvc = hw
	return hs
}

// Run implements Run api server function for Http server
func (hs *HttpServer) Run(ctx context.Context, cancelFn context.CancelFunc) error {
	var err error = nil
	go func() {
		if err := hs.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hs.Logger.Fatal("http server stopped with error", zap.Error(err))
		}
	}()

	hs.Logger.Info("app started")

	<-ctx.Done()
	cancelFn()
	hs.Logger.Info("shutting down gracefully start")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = hs.Server.Shutdown(ctxTimeout); err != nil {
		hs.Logger.Error("error shutting server down", err)
	}
	return err
}
