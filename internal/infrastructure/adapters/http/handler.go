package http

import (
	"context"
	"net/http"

	"github.com/beabys/go-template/internal/app/config"
	"github.com/beabys/go-template/internal/application/example/handler"
	"github.com/beabys/go-template/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type HttpServer struct {
	Server         *http.Server
	Config         *config.Config
	Logger         logger.Logger
	ExampleService handler.ExampleServiceHandler
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (hs *HttpServer) SetConfig(c *config.Config) *HttpServer {
	hs.Config = c
	return hs
}

func (hs *HttpServer) SetLogger(l logger.Logger) *HttpServer {
	hs.Logger = l
	return hs
}

func (hs *HttpServer) SetExampleService(svc handler.ExampleServiceHandler) *HttpServer {
	hs.ExampleService = svc
	return hs
}

func (hs *HttpServer) Run(ctx context.Context, wg *errgroup.Group) {
	wg.Go(func() error {
		hs.Logger.Info("http server started")
		if err := hs.Server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				return nil
			}
			hs.Logger.Error("http server stopped with error", err)
			return err
		}
		return nil
	})

	wg.Go(func() error {
		<-ctx.Done()
		hs.Logger.Info("shutting down gracefully http server")
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5)
		defer cancel()
		if err := hs.Server.Shutdown(ctxTimeout); err != nil {
			hs.Logger.Error("error shutting server down", err)
			return err
		}
		return nil
	})
}
