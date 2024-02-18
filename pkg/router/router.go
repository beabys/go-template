package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
	"gitlab.com/beabys/quetzal"
)

type Router interface {
	GetMuxHandler() interface{}
	SetDefaultMiddlewares()
	Serve(string, int)
}

type DefaultRouter struct {
	Mux    *chi.Mux
	logger quetzal.Logger
}

// NewRouter return Router Type
func NewDefaultRouter() *DefaultRouter {
	logger := quetzal.NewDefaultLogger(&quetzal.DefaultLoggerConfig{})
	return &DefaultRouter{
		Mux:    chi.NewRouter(),
		logger: logger,
	}
}

func (r *DefaultRouter) GetMuxHandler() interface{} {
	return r.Mux
}

func (r *DefaultRouter) SetLogger(l quetzal.Logger) *DefaultRouter {
	r.logger = l
	return r
}

func (r *DefaultRouter) SetDefaultMiddlewares() {
	r.Mux.Use(m.RealIP)
	r.Mux.Use(Recoverer(r.logger))
	r.Mux.Use(JsonContentType)
	r.Mux.NotFound(NotFound)
}

// Serve web application on port
func (r *DefaultRouter) Serve(host string, port int) {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%v", host, port),
		Handler: r.Mux,
	}

	go func() {
		// signChan channel is used to transmit signal notifications.
		signChan := make(chan os.Signal, 1)
		// Catch and relay certain signal(s) to signChan channel.
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
		// Blocking until a signal is sent over signChan channel.
		sig := <-signChan

		r.logger.Info(fmt.Sprintf("router - shutdown: %v", sig))

		// Create a new context with a timeout duration. It helps allowing
		// timeout to all active connections in order  to complete their job.
		// Any connections that wont complete within the allowed timeout gets halted.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err == context.DeadlineExceeded {
			r.logger.Info("router - shutdown: halted active connections")
		}
	}()

	// Starting Server
	r.logger.Info(fmt.Sprintf("http server listening on %v", server.Addr))

	err := server.ListenAndServe()

	switch true {
	case err == http.ErrServerClosed:
		r.logger.Info("router - shutdown: started")
	case err != nil:
		r.logger.Info(err)
	}
}
