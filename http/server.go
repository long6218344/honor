package http

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*Context)

// Server http server.
type Server struct {
	opts   *options
	engine *gin.Engine
}

// New returns a http server.
func New(opts ...Option) *Server {
	srv := new(Server)
	for _, opt := range opts {
		opt(srv.opts)
	}
	srv.engine = gin.New()
	return srv
}

// Start start the http server.
func (srv *Server) Start() error {
	if _, _, err := net.SplitHostPort(srv.opts.addr); err != nil {
		panic(err)
	}
	go func() {
		if err := http.ListenAndServe(srv.opts.addr, srv.engine); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (srv *Server) Use(middleware ...HandlerFunc) {
	var ms []gin.HandlerFunc
	for _, m := range middleware {
		ms = append(ms, func(ctx *gin.Context) {
			m(&Context{ctx})
		})
	}
	srv.engine.Use(ms...)
}

func (srv *Server) Stop() {
}
