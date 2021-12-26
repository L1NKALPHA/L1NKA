package gmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus implement
type Prometheus struct {
}

func NewPrometheus() IMetrics {
	return &Prometheus{}
}

// promCounter counter vec.
type promCounter struct {
	counter *prometheus.CounterVec
}

func (*Prometheus) Counter(cfg *CounterOpts) ICounter {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
		}, cfg.Labels)
	prometheus.MustRegister(vec)
	return &promCounter{
		counter: vec,
	}
}

// Inc increments the counter by 1. Use Add to increment it by arbitrary.
func (counter *promCounter) Inc(labels ...string) {
	counter.counter.WithLabelValues(labels...).Inc()
}

// Add increments the counter by 1. Use Add to increment it by arbitrary.
func (counter *promCounter) Add(v float64, labels ...string) {
	counter.counter.WithLabelValues(labels...).Add(v)
}

// gaugeVec gauge vec.
type promGaugeVec struct {
	gauge *prometheus.GaugeVec
}

// Gauge .
func (*Prometheus) Gauge(cfg *GaugeOpts) IGauge {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
		}, cfg.Labels)
	prometheus.MustRegister(vec)
	return &promGaugeVec{
		gauge: vec,
	}
}

// Inc increments the counter by 1. Use Add to increment it by arbitrary.
func (gauge *promGaugeVec) Inc(labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Inc()
}

// Add increments the counter by 1. Use Add to increment it by arbitrary.
func (gauge *promGaugeVec) Add(v float64, labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Add(v)
}

// Set the given value to the collection.
func (gauge *promGaugeVec) Set(v float64, labels ...string) {
	gauge.gauge.WithLabelValues(labels...).Set(v)
}

// Histogram prom histogram collection.
type promHistogram struct {
	histogram *prometheus.HistogramVec
}

// Histogram new a histogram vec.
func (*Prometheus) Histogram(cfg *HistogramOpts) IHistogram {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
			Buckets:   cfg.Buckets,
		}, cfg.Labels)
	prometheus.MustRegister(vec)
	return &promHistogram{
		histogram: vec,
	}
}

// Observe Timing adds a single observation to the histogram.
func (histogram *promHistogram) Observe(v int64, labels ...string) {
	histogram.histogram.WithLabelValues(labels...).Observe(float64(v))
}
