package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_reqs",
		Help: "Total number of gRPC requests",
	})

	HTTPResponsesMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_responses",
		Help: "Total number of http requests",
	}, []string{"code", "endpoint"})
)

func SetupMetrics() {
	prometheus.Register(TotalRequestsCounter)
	prometheus.Register(HTTPResponsesMetric)
}
