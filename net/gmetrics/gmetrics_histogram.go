package gmetrics

// HistogramOpts is histogram vector opts.
type HistogramOpts struct {
	Opts
	Buckets []float64
}

// IHistogram gauge vec.
type IHistogram interface {
	// Observe adds a single observation to the histogram.
	Observe(v int64, labels ...string)
}
