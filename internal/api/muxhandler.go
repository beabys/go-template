package api

import (
	"errors"
	"fmt"
	"net/http"

	v1 "github.com/beabys/go-template/internal/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMuxHandler(server *HttpServer) (http.Handler, error) {
	swagger, err := v1.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error loading swagger spec\n: %w", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	r := chi.NewRouter() // http.Handler

	// metrics latency
	prometheusMetrics := NewPrometheusMetrics()
	prometheus.MustRegister(prometheusMetrics.latency)

	// default middlewars
	r.NotFound(notFound)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middlewareMetrics(server.Logger, prometheusMetrics))

	// public auth router group
	r.Group(func(r chi.Router) {
		r.Mount("/metrics", promhttp.Handler())
	})

	r.Group(func(r chi.Router) {
		// cors
		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders: []string{
				"Accept", "Authorization", "Content-Type", "X-CSRF-Token",
				"Access-Control-Allow-Headers", "X-Requested-With",
				"Access-Control-Request-Method", "Access-Control-Request-Headers",
			},
			MaxAge: 300, // Maximum value not ignored by any of major browsers
		}))

		// Mount oapi routes
		v1.HandlerWithOptions(server, v1.ChiServerOptions{
			BaseRouter: r,
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				errorResponseJSON(w, http.StatusInternalServerError, err)
			},
		})

	})

	return r, nil
}

// Not found Middleware
func notFound(w http.ResponseWriter, r *http.Request) {
	errorResponseJSON(w, http.StatusNotFound, errors.New("not found"))
}

func DefaultError(w http.ResponseWriter, r *http.Request, err error) {
	errorResponseJSON(w, http.StatusInternalServerError, err)
}

func JsonContentType(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
