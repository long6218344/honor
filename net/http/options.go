package http

import "time"

type options struct {
	addr    string
	timeout time.Duration
}

type Option func(opts *options)

func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}
