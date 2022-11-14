package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var pingProcessingTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "ping_request_processing_time",
		Buckets: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
	}, []string{"Dimension1", "Dimension2"},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	start := time.Now()
	fmt.Fprintf(w, "pong")
	randNum := rand.Intn(5)
	time.Sleep(time.Duration(randNum) * time.Second)
	pingProcessingTime.With(map[string]string{"Dimension1": "1", "Dimension2": "2"}).Observe(float64(time.Since(start).Seconds()))
}

func main() {
	prometheus.MustRegister(pingCounter, pingProcessingTime)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8090", nil)
}
