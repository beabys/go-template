package api

import (
	"net/http"
	"time"

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

func middlewareLatency(p *PrometheusMetrics) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			elapsed := time.Since(startTime).Seconds()
			p.latency.WithLabelValues(
				r.Method,
				r.URL.Path,
			).Observe(elapsed)
		}
		return http.HandlerFunc(fn)
	}
}
