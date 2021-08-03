package main

import (
	"strconv"
	"time"

	"github.com/lacazethomas/nicehash-exporter/config"
	"github.com/lacazethomas/nicehash-exporter/models"
	"github.com/lacazethomas/nicehash-exporter/route"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics listed are supported
type Metrics struct {
	temperatureVRAM     *prometheus.GaugeVec
	temperatureDevice   *prometheus.GaugeVec
	miningSpeed         *prometheus.GaugeVec
	walletBalance       prometheus.Gauge
	unpaidAmount        prometheus.Gauge
	nextpayouttimestamp prometheus.Gauge
	feerules            prometheus.Gauge
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

	metrics.nextpayouttimestamp = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "nextpayouttimestamp",
		Help:      "next payout timestamp",
	})

	metrics.feerules = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "nicehash",
		Name:      "feerules",
		Help:      "feerules in pourcent",
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

			m.nextpayouttimestamp.Set(float64(g.Nextpayouttimestamp.Unix()))

			f, err := route.GetFeeRules(appConfig)
			Check("Getting FeeRules metrics", err)
			m.feerules.Set(float64(f.Withdrawal.External.Rules.Btc.Intervals[0].Element.Value))

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

					if len(g.Miningrigs[i].Devices[y].Speeds) == 0 {
						m.miningSpeed.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name).Set(0)
					} else {
						for z := range g.Miningrigs[i].Devices[y].Speeds {
							speedF, err := strconv.ParseFloat(g.Miningrigs[i].Devices[y].Speeds[z].Speed, 64)
							if err == nil {
								m.miningSpeed.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name, g.Miningrigs[i].Devices[y].Speeds[z].Algorithm).Set(speedF)
							} else {
								m.miningSpeed.WithLabelValues(g.Miningrigs[i].Name, g.Miningrigs[i].Devices[y].Name, g.Miningrigs[i].Devices[y].Speeds[z].Algorithm).Set(0)
							}
						}
					}

				}

			}
			time.Sleep(60 * time.Second)
		}
	}()
}
