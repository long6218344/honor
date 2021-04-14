package rpc

import (
	"math"
	"net"
	"os"

	"google.golang.org/grpc"
)

const maxInterceptors int8 = math.MaxInt8 / 2

// Server rpc server.
type Server struct {
	opts *options

	srv          *grpc.Server
	interceptors []grpc.UnaryServerInterceptor
}

// New .
func New(opts ...Option) *Server {
	srv := new(Server)
	for _, opt := range opts {
		opt(srv.opts)
	}
	srv.srv = grpc.NewServer()
	return srv
}

// User .
func (srv *Server) Use(interceptors ...grpc.UnaryServerInterceptor) {
	srv.interceptors = srv.combineInterceptors(interceptors...)
}

// Srv .
func (srv *Server) Srv() *grpc.Server {
	return srv.srv
}

func (srv *Server) combineInterceptors(interceptors ...grpc.UnaryServerInterceptor) []grpc.UnaryServerInterceptor {
	finalSize := len(srv.interceptors) + len(interceptors)
	if finalSize >= int(maxInterceptors) {
		panic("too many interceptors")
	}
	mergedInterceptors := make([]grpc.UnaryServerInterceptor, finalSize)
	copy(mergedInterceptors, srv.interceptors)
	copy(mergedInterceptors[len(srv.interceptors):], interceptors)
	return mergedInterceptors
}

func (srv *Server) Run(addrs ...string) error {
	addr := resolveAddress(addrs)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.srv.Serve(listen)
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
