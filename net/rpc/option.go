package rpc

import (
	"google.golang.org/grpc"
	"time"
)

const (
	defaultTime              = time.Hour
	defaultTimeout           = time.Second * 20
	defaultMaxConnectionAge  = time.Hour
	defaultMaxConnectionIdle = time.Hour
)

type options struct {
	time              time.Duration
	timeout           time.Duration
	maxConnectionIdle time.Duration
	maxConnectionAge  time.Duration

	// interceptor
	unaryInt       grpc.UnaryServerInterceptor
	chainUnaryInts []grpc.UnaryServerInterceptor
}

var defaultOptions = options{
	time:              defaultTime,
	timeout:           defaultTimeout,
	maxConnectionIdle: defaultMaxConnectionIdle,
	maxConnectionAge:  defaultMaxConnectionAge,
}

type Option func(*options)

// WithTimeout .
func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

// WithUnaryInterceptor .
func WithUnaryInterceptor(int grpc.UnaryServerInterceptor) Option {
	return func(opts *options) {
		opts.unaryInt = int
	}
}

// WithChainUnaryInterceptors .
func WithChainUnaryInterceptors(ints []grpc.UnaryServerInterceptor) Option {
	return func(opts *options) {
		opts.chainUnaryInts = append(opts.chainUnaryInts, ints...)
	}
}
