package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

func Init(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Info().Msgf("metrics server listening on %s", addr)
		// Start the HTTP server for metric
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Error().Err(err).Msg("metrics server failed")
		}
	}()
	log.Info().Msgf("metrics server listening on %s", addr)
	log.Info().Msg("metrics server started")
	// Register custom metrics
	prometheus.MustRegister(TodoCreateCounter)
}

var TodoCreateCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "todos_created_total",
		Help: "Total number of todos created, labeled by status",
	},
	[]string{"status"},
)
