package ghttp

import (
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/net/gmetrics"
)

var (
	metricServerReqDur = gmetrics.Histogram(&gmetrics.HistogramOpts{
		Opts: gmetrics.Opts{
			Namespace: "http_server",
			Subsystem: "requests",
			Name:      "duration_ms",
			Help:      "http server requests duration(ms).",
			Labels:    []string{"path", "caller"},
		},
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})
	metricServerReqCodeTotal = gmetrics.Counter(&gmetrics.CounterOpts{
		Namespace: "http_server",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "http server requests error count.",
		Labels:    []string{"path", "caller", "code"},
	})
)

// Metrics default metrics middleware
func Metrics(r *Request) {
	now := time.Now()
	path := r.URL.Path
	r.Middleware.Next()
	// caller is which service is calling you

	caller := strings.ToUpper(r.GetHeader("X-Caller"))
	if len(caller) == 0 {
		caller = "OTHERS"
	}
	rt := time.Since(now)

	if len(path) > 0 {
		metricServerReqCodeTotal.Inc(path, caller, gvar.New(r.Response.Status).String())
		metricServerReqDur.Observe(int64(rt/time.Millisecond), path, caller)
	}
}
