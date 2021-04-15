package rpc

import (
	"context"
	"google.golang.org/grpc/reflection"
	"math"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const maxInterceptors int8 = math.MaxInt8 / 2

// Server rpc server.
type Server struct {
	opts options
	srv  *grpc.Server
}

// New create a rpc server instance.
func New(opt ...Option) *Server {
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	srv := &Server{
		opts: opts,
	}
	srv.srv = grpc.NewServer(srv.serverOptions()...)
	return srv
}

// get grpc server options.
func (s *Server) serverOptions() (opts []grpc.ServerOption) {
	param := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:              s.opts.time,
		Timeout:           s.opts.timeout,
		MaxConnectionAge:  s.opts.maxConnectionAge,
		MaxConnectionIdle: s.opts.maxConnectionIdle,
	})
	opts = append(opts, param, grpc.UnaryInterceptor(s.chainUnaryInterceptors))
	return
}

func (s *Server) chainUnaryInterceptors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	interceptors := s.opts.chainUnaryInts
	if s.opts.unaryInt != nil {
		interceptors = append([]grpc.UnaryServerInterceptor{s.opts.unaryInt}, s.opts.chainUnaryInts...)
	}
	if len(interceptors) == 0 {
		return handler(ctx, req)
	} else if len(interceptors) == 1 {
		return interceptors[0](ctx, req, info, handler)
	}
	return interceptors[0](ctx, req, info, getChainUnaryHandler(interceptors, 0, info, handler))
}

func getChainUnaryHandler(interceptors []grpc.UnaryServerInterceptor, idx int, info *grpc.UnaryServerInfo, finalHandler grpc.UnaryHandler) grpc.UnaryHandler {
	if len(interceptors)-1 == idx {
		return finalHandler
	}
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptors[idx+1](ctx, req, info, getChainUnaryHandler(interceptors, idx+1, info, finalHandler))
	}
}

// Use add interceptors.
func (s *Server) Use(interceptors ...grpc.UnaryServerInterceptor) {
	s.opts.chainUnaryInts = s.combineInterceptors(interceptors...)
}

func (s *Server) combineInterceptors(interceptors ...grpc.UnaryServerInterceptor) []grpc.UnaryServerInterceptor {
	finalSize := len(s.opts.chainUnaryInts) + len(interceptors)
	if finalSize >= int(maxInterceptors) {
		panic("too many interceptors")
	}
	mergedInterceptors := make([]grpc.UnaryServerInterceptor, finalSize)
	copy(mergedInterceptors, s.opts.chainUnaryInts)
	copy(mergedInterceptors[len(s.opts.chainUnaryInts):], interceptors)
	return mergedInterceptors
}

// Server get grpc server.
func (s *Server) Server() *grpc.Server {
	return s.srv
}

// Run start rpc server.
func (s *Server) Run(addrs ...string) error {
	addr := resolveAddress(addrs)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	reflection.Register(s.srv)
	return s.srv.Serve(listen)
}

func (s *Server) Close() {
	s.srv.GracefulStop()
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("RPC_PORT"); port != "" {
			return ":" + port
		}
		return ":9090"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
