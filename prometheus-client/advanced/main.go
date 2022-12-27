package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var wave = prometheus.NewGauge(prometheus.GaugeOpts{
	Name:        "wave",
	Help:        "sinus function current value",
	ConstLabels: prometheus.Labels{"label": "mylabel"},
})

var degrees float64 = 0.0

func runWaveGenerator() {
	for {
		value := math.Sin(degrees * math.Pi / 180.0)
		wave.Set(value)
		degrees += 15.0
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Running metrics provider at localhost:8080/metrics")

	go runWaveGenerator()
	prometheus.MustRegister(wave)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
