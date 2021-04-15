package serverinterceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Limiter defines the interface to perform request rate limiting.
// If Limit return true, the request will be rejexted.
// Otherwise, the request will pass.
type Limiter interface {
	Limit() bool
}

func RateLimitInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if limiter.Limit() {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by rate limit", info.FullMethod)
		}
		return handler(ctx, req)
	}
}
