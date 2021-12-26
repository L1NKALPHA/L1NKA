package gredis

import (
	"github.com/gogf/gf/v2/net/gmetrics"
	"github.com/gogf/gf/v2/text/gstr"
)

const namespace = "redis_client"

var (
	metricReqDur = gmetrics.Histogram(&gmetrics.HistogramOpts{
		Opts: gmetrics.Opts{
			Namespace: namespace,
			Subsystem: "requests",
			Name:      "duration_ms",
			Help:      "redis client requests duration(ms).",
			Labels:    []string{"name", "addr", "command"},
		},
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2500},
	})
	metricReqErr = gmetrics.Counter(&gmetrics.CounterOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "redis client requests error count.",
		Labels:    []string{"name", "addr", "command", "error"},
	})
	metricHits = gmetrics.Counter(&gmetrics.CounterOpts{
		Namespace: namespace,
		Subsystem: "",
		Name:      "hits_total",
		Help:      "redis client hits total.",
		Labels:    []string{"name", "addr"},
	})
	metricMisses = gmetrics.Counter(&gmetrics.CounterOpts{
		Namespace: namespace,
		Subsystem: "",
		Name:      "misses_total",
		Help:      "redis client misses total.",
		Labels:    []string{"name", "addr"},
	})
)

func formatErr(err error) string {
	es := err.Error()
	switch {
	case gstr.Contains(es, "read"):
		return "read timeout"
	case gstr.Contains(es, "dial"):
		return "dial timeout"
	case gstr.Contains(es, "write"):
		return "write timeout"
	case gstr.Contains(es, "EOF"):
		return "eof"
	case gstr.Contains(es, "reset"):
		return "reset"
	case gstr.Contains(es, "broken"):
		return "broken pipe"
	default:
		return "unexpected err"
	}
}
