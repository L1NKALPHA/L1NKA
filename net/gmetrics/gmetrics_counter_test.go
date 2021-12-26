package gmetrics

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

var (
	prometheusClient = NewPrometheus()
)

func Test_PrometheusCounterVec(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		counterVec := prometheusClient.Counter(&CounterOpts{
			Namespace: "test",
			Subsystem: "test",
			Name:      "test",
			Help:      "this is test metrics.",
			Labels:    []string{"name", "addr"},
		})
		counterVec.Inc("name1", "127.0.0.1")

		t.Panics(func() {
			prometheusClient.Counter(&CounterOpts{
				Namespace: "test",
				Subsystem: "test",
				Name:      "test",
				Help:      "this is test metrics.",
				Labels:    []string{"name", "addr"},
			})
		})
		t.NotPanics(func() {
			prometheusClient.Counter(&CounterOpts{
				Namespace: "test",
				Subsystem: "test",
				Name:      "test2",
				Help:      "this is test metrics.",
				Labels:    []string{"name", "addr"},
			})
		})
	})
}
