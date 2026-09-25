package server

import (
	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/handlers"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/store"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/webhook"
	"github.com/gin-contrib/cors"
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
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.WebPortalFrontURL},
		AllowMethods: []string{"GET", "POST"},
	}))

	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.POST("/webhook/release", webhook.HandleGitRelease(s.store))
	api := s.router.Group("/api")
	{
		api.GET("/platform-tools-services", handlers.ApiPlatformToolsServices(s.cfg))
		api.GET("/git-releases", handlers.ApiGitReleases(s.store))
		api.GET("/git-repositories", handlers.ApiGitRepositories(s.cfg))
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
