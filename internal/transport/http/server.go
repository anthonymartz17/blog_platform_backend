package http

import (
	"context"
	"errors"
	"net/http"
)

// ShutdownFunc runs cleanup logic during server shutdown.
type ShutdownFunc func(context.Context) error

// Server wraps an http.Server
type Server struct {
	addr        string
	httpServer  *http.Server
	shutdownFns []ShutdownFunc
}

// New creates a new Server
func NewServer(addr string, router http.Handler, shutdownFns ...ShutdownFunc) *Server {
	return &Server{
		addr: addr,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		shutdownFns: shutdownFns,
	}
}

// Start starts the Server and listens for requets
func (s *Server) Start() error {

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully stops the HTTP server and releases managed resources.
func (s *Server) Shutdown(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	for _, shutdownFn := range s.shutdownFns {
		if shutdownErr := shutdownFn(ctx); shutdownErr != nil && err == nil {
			err = shutdownErr
		}
	}

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (s *Server) Address() string {
	return s.addr
}
