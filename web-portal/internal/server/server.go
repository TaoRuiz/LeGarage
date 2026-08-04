package server

import (
	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
}

func New(cfg *config.Config) *Server {
	s := &Server{
		router: gin.Default(),
		cfg:    cfg,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.Static("/static", "./static")
	s.router.LoadHTMLGlob("templates/*")
	s.router.GET("/", handlers.Index(s.cfg))
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
