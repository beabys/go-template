package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "gitlab.com/beabys/go-http-template/api/v1"
	"gitlab.com/beabys/go-http-template/internal/api"
	"gitlab.com/beabys/quetzal"
)

const (
	apiPrefix = "/v1"
)

func NewMuxHandler(ctx context.Context, server *api.HttpServer) http.Handler {
	r := chi.NewRouter() // http.Handler

	// public auth router group
	r.Group(func(r chi.Router) {
		r.Mount("/metrics", promhttp.Handler())
	})

	r.NotFound(NotFound)

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

		r.Use(JsonContentType)
		r.Use(middleware.Recoverer)

		r.Mount(apiPrefix, v1.HandlerWithOptions(server, v1.ChiServerOptions{
			BaseRouter:       r,
			BaseURL:          apiPrefix,
			ErrorHandlerFunc: DefaultError,
		}))

	})

	return r
}

func DefaultError(w http.ResponseWriter, r *http.Request, err error) {
	msg := map[string]interface{}{
		"errot": err.Error(),
	}
	response := &v1.Response{
		Error: &msg,
	}
	quetzal.ResponseJSON(w, http.StatusInternalServerError, response)
}

func JsonContentType(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// Not found Middleware
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}