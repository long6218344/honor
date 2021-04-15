package serverinterceptors

import (
	"context"

	"google.golang.org/grpc"
)

func RecoverInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		return handler(ctx, req)
	}
}
