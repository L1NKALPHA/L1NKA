package gdb

import (
	"github.com/gogf/gf/v2/net/gmetrics"
)

const namespace = "sql_client"

var (
	metricReqDur = gmetrics.Histogram(&gmetrics.HistogramOpts{
		Opts: gmetrics.Opts{
			Namespace: namespace,
			Subsystem: "requests",
			Name:      "duration_ms",
			Help:      "redis client requests duration(ms).",
			Labels:    []string{"name", "addr", "schema"},
		},
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2500},
	})
	metricReqErr = gmetrics.Counter(&gmetrics.CounterOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "redis client requests error count.",
		Labels:    []string{"name", "addr", "schema", "error"},
	})
)
