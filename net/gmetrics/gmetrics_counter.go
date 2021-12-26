package gmetrics

// CounterOpts is an alias of Opts.
type CounterOpts Opts

// ICounter counter vec.
type ICounter interface {
	// Inc increments the counter by 1. Use Add to increment it by arbitrary
	// non-negative values.
	Inc(labels ...string)
	// Add adds the given value to the counter. It panics if the value is <
	// 0.
	Add(v float64, labels ...string)
}
