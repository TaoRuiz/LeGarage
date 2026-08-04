package server

import (
	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/handlers"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/store"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/webhook"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	store  *store.Store
}

func New(cfg *config.Config, store *store.Store) *Server {
	s := &Server{
		router: gin.Default(),
		cfg:    cfg,
		store:  store,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.Static("/static", "./static")
	s.router.LoadHTMLGlob("templates/*")
	s.router.GET("/", handlers.Index(s.cfg, s.store))
	s.router.POST("/webhook/release", webhook.HandleRelease(s.store))
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
