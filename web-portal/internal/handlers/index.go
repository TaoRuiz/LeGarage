package handlers

import (
	"net/http"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
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

	return func(c *gin.Context) {
		statuses := make([]ServiceStatus, len(services))
		for i, s := range services {
			statuses[i] = ServiceStatus{
				Name:    s.Name,
				Healthy: checkService(s.URL),
			}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Services": statuses,
		})
	}
}
