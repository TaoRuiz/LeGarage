package server

import (
	"net/http"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/handlers"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/", handlers.Index)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
