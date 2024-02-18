package router

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"gitlab.com/beabys/quetzal"
)

// recoverer catch panics and return 500 error, logging the error and trace stack
func Recoverer(logger quetzal.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					stackTrace := fmt.Sprintf("%v\n%v", rec, string(debug.Stack()))
					message := "Recovering from panic"
					logger.Error(message, fmt.Errorf(stackTrace))
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
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
