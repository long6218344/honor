package rpc

import "time"

type options struct {
	Timeout time.Duration
}

type Option func(*options)

// WithTimeout .
func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.Timeout = timeout
	}
}
