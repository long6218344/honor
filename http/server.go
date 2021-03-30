package http

// Server http server.
type Server struct {
	opts *options
}

// New returns a http server.
func New(opts ...Option) *Server {
	srv := new(Server)
	for _, opt := range opts {
		opt(srv.opts)
	}
	return srv
}

func (srv *Server) Start() error {
	return nil
}

func (srv *Server) Stop() {

}
