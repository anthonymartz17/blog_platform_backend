package http

import "net/http"

//Server wraps an http.Server
type Server struct{
	addr string
	httpServer *http.Server
}

//New creates a new Server
func NewServer(addr string, router http.Handler) *Server{
   return &Server{
	   addr: addr,
		 httpServer: &http.Server{
			Addr: addr,
			Handler: router,
		 },
	 }
}

//Start starts the Server and listens for requets
func (s *Server) Start() error{
	
	return s.httpServer.ListenAndServe()
}

func (s *Server) Address() string{
   return s.addr
}