package main

import (
	"strconv"
	"time"

	"github.com/lacazethomas/nicehashExporter/config"
	"github.com/lacazethomas/nicehashExporter/models"
	"github.com/lacazethomas/nicehashExporter/route"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics listed are supported
type Metrics struct {
	temperatureVRAM   *prometheus.GaugeVec
	temperatureDevice *prometheus.GaugeVec
	miningSpeed       *prometheus.GaugeVec
	walletBalance     prometheus.Gauge
	unpaidAmount      prometheus.Gauge
}

func initMetrics() (metrics Metrics) {

	metrics = Metrics{}
	metrics.temperatureVRAM = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "temperatureVRAM",
		Help:      "Temperature in °C",
	}, []string{"localisation", "device"})

	metrics.temperatureDevice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "temperatureDevice",
		Help:      "Temperature in °C",
	}, []string{"localisation", "device"})

	metrics.miningSpeed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "miningSpeed",
		Help:      "Mining speed in MH",
	}, []string{"localisation", "device", "algo"})

	metrics.walletBalance = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "walletbalance",
		Help:      "Balance in BTC",
	})

	metrics.unpaidAmount = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "unpaidAmount",
		Help:      "unpaid in BTC",
	})

	return
}

func getMetrics(appConfig config.AppConfig, m Metrics) {
	go func() {
		for {

			b, err := route.GetBalance(appConfig)
			Check("Getting balance metrics", err)

			availableF, err := strconv.ParseFloat(b.Available, 64)
			Check("Parsing Available string to float64", err)
			if err == nil {
				m.walletBalance.Set(float64(availableF))
			}

			g, err := route.GetMiningStats(appConfig)
			Check("Getting miningStats metrics", err)

			pendingF, err := strconv.ParseFloat(g.Unpaidamount, 64)
			Check("Parsing Unpaidamount string to float64", err)
			if err == nil {
				m.unpaidAmount.Set(float64(pendingF))
			}

			g.ComputeTemperature()

			for i := range g.Miningrigs {
				for y := range g.Miningrigs[i].Devices {
					if g.Miningrigs[i].Devices[y].Temperature > 0 {
						if !(&models.Temperature{T: g.Miningrigs[i].Devices[y].Temperature}).IsRealTemp() {
							m.temperatureDevice.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name).Set(float64(g.Miningrigs[i].Devices[y].GPUTemperature))
							m.temperatureVRAM.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name).Set(float64(g.Miningrigs[i].Devices[y].VRAMTemperature))
						} else {
							m.temperatureDevice.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name).Set(g.Miningrigs[i].Devices[y].Temperature)
						}
					}

					for z := range g.Miningrigs[i].Devices[y].Speeds {
						speedF, err := strconv.ParseFloat(g.Miningrigs[i].Devices[y].Speeds[z].Speed, 64)
						Check("Parsing Speed string to float64", err)
						if err == nil {
							m.miningSpeed.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name, g.Miningrigs[i].Devices[y].Speeds[z].Algorithm).Set(speedF)
						}
					}

				}

			}
			time.Sleep(60 * time.Second)
		}
	}()
}
