package gclient

import (
	"net/http"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/net/gmetrics"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	metricClientReqDur = gmetrics.Histogram(&gmetrics.HistogramOpts{
		Opts: gmetrics.Opts{
			Namespace: "http_client",
			Subsystem: "requests",
			Name:      "duration_ms",
			Help:      "http client requests duration(ms).",
			Labels:    []string{"path", "caller"},
		},
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})
	metricClientReqCodeTotal = gmetrics.Counter(&gmetrics.CounterOpts{
		Namespace: "http_client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "http client requests code count.",
		Labels:    []string{"path", "caller", "code"},
	})
)

func (c *Client) SetCaller(caller string) *Client {
	return c.SetHeader("X-Caller", caller)
}

// Metrics default metrics middleware
func Metrics(c *Client, r *http.Request) (response *Response, err error) {
	now := time.Now()
	path := r.URL.Path

	response, err = c.Next(r)
	// caller is who you are calling
	caller := gconv.String(r.Context().Value("X-Caller"))
	rt := time.Since(now)
	statusCode := 200
	if response != nil {
		statusCode = response.StatusCode
	} else {
		statusCode = 500
	}

	if len(path) > 0 {
		metricClientReqCodeTotal.Inc(path, caller, gvar.New(statusCode).String())
		metricClientReqDur.Observe(int64(rt/time.Millisecond), path, caller)
	}
	return
}
