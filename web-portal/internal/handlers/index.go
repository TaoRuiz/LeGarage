package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/forgejo"
	"github.com/gin-gonic/gin"
)

type RepoWithCommits struct {
	Repo    forgejo.Repo
	Commits []forgejo.Commit
}

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

		// Commits
		reposWithCommits := make([]RepoWithCommits, 0, len(repos))
		for _, repo := range repos {
			parts := strings.Split(repo.FullName, "/")
			if len(parts) != 2 {
				continue
			}
			commits, err := forgejoClient.GetCommits(parts[0], parts[1], 3)
			if err != nil {
				log.Printf("commits error for %s: %v", repo.FullName, err)
			}
			reposWithCommits = append(reposWithCommits, RepoWithCommits{
				Repo:    repo,
				Commits: commits,
			})
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Services": statuses,
			"Repos":    reposWithCommits,
		})
	}
}
