package api

import (
	"net/http"
	"time"

	"github.com/beabys/go-template/pkg/logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

func NewPrometheusMetrics() *PrometheusMetrics {
	latency := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "latency_duration_seconds",
			Help:       "latency duration distribution",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.01},
		},
		[]string{"method", "path"},
	)
	return &PrometheusMetrics{latency}
}

func middlewareMetrics(log logger.Logger, p *PrometheusMetrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()
			defer func() {
				elapsed := time.Since(start)
				p.latency.WithLabelValues(
					r.Method,
					r.URL.Path,
				).Observe(elapsed.Seconds())
				log.Info(
					"request",
					logger.LogField{Key: "request-id", Value: r.Header.Get(middleware.RequestIDHeader)},
					logger.LogField{Key: "method", Value: r.Method},
					logger.LogField{Key: "path", Value: r.URL.Path},
					logger.LogField{Key: "query", Value: r.URL.RawQuery},
					logger.LogField{Key: "ip", Value: r.RemoteAddr},
					logger.LogField{Key: "user-agent", Value: r.UserAgent()},
					logger.LogField{Key: "status", Value: ww.Status()},
					logger.LogField{Key: "latency", Value: elapsed},
					logger.LogField{Key: "bytes", Value: ww.BytesWritten()},
					logger.LogField{Key: "latency", Value: elapsed},
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
