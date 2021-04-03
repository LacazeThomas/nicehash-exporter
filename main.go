package main

import (
	"log"
	"net/http"

	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	var err error

	appConfig := config.GetAppConfig()

	// Set log level
	if appConfig.IsDev() {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Println("Error to initialize logger")
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	metrics := initMetrics()
	prometheus.Register(metrics.temperatureDevice)
	prometheus.Register(metrics.temperatureVRAM)
	prometheus.Register(metrics.walletBalance)
	prometheus.Register(metrics.unpaidAmount)
	prometheus.Register(metrics.miningSpeed)

	getMetrics(appConfig, metrics)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9159", nil)
}
