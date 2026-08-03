package handlers

import (
	"log"
	"net/http"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/forgejo"
	"github.com/gin-gonic/gin"
)

func Index(cfg *config.Config) gin.HandlerFunc {
	services := []struct {
		Name string
		URL  string
	}{
		{"Forgejo", cfg.ForgejoURL},
		{"Authentik", cfg.AuthentikURL},
		{"Grafana", cfg.GrafanaURL},
		{"Prometheus", cfg.PrometheusURL},
	}

	forgejoClient := forgejo.NewClient(cfg.ForgejoURL)

	return func(c *gin.Context) {
		// Services statuses
		statuses := make([]ServiceStatus, len(services))
		for i, s := range services {
			statuses[i] = ServiceStatus{
				Name:    s.Name,
				Healthy: checkService(s.URL),
			}
		}

		// Repositories
		repos, err := forgejoClient.GetPublicRepos()
		if err != nil {
			log.Printf("forgejo error: %v", err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Services": statuses,
			"Repos":    repos,
		})
	}
}
