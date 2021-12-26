package gmetrics

// Opts contains the common arguments for creating vec Metric..
type Opts struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
}

var (
	defaultMetrics IMetrics
)

func init() {
	defaultMetrics = NewPrometheus()
}

type IMetrics interface {
	Counter(cfg *CounterOpts) ICounter
	Gauge(cfg *GaugeOpts) IGauge
	Histogram(cfg *HistogramOpts) IHistogram
}

func Counter(cfg *CounterOpts) ICounter {
	return defaultMetrics.Counter(cfg)
}

func Gauge(cfg *GaugeOpts) IGauge {
	return defaultMetrics.Gauge(cfg)
}

func Histogram(cfg *HistogramOpts) IHistogram {
	return defaultMetrics.Histogram(cfg)
}
