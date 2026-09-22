package handlers

import (
	"net/http"
	"strings"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/config"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/dto"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/forgejo"
	"github.com/TaoRuiz/LeGarage/web-portal/internal/store"
	"github.com/gin-gonic/gin"
)

func ApiPlatformToolsServices(cfg *config.Config) gin.HandlerFunc {
	services := []platformToolsService{
		{Name: "Forgejo", URL: cfg.ForgejoURL},
		{Name: "Authentik", URL: cfg.AuthentikURL},
		{Name: "Grafana", URL: cfg.GrafanaURL},
		{Name: "Prometheus", URL: cfg.PrometheusURL},
	}

	return func(c *gin.Context) {
		statuses := make([]dto.PlatformToolsServiceResponse, len(services))
		for i, s := range services {
			statuses[i] = dto.PlatformToolsServiceResponse{
				Name:    s.Name,
				Healthy: checkService(s.URL),
			}
		}
		c.JSON(http.StatusOK, statuses)
	}
}

func ApiGitReleases(s *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		releases, err := s.ListReleases(20)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		result := make([]dto.GitReleaseResponse, len(releases))
		for i, r := range releases {
			result[i] = dto.GitReleaseResponse{
				ID:             r.ID,
				RepositoryName: r.Repo,
				Tag:            r.TagName,
				TagName:        r.Name,
				Author:         r.Author,
				URL:            r.HTMLURL,
				PublishedAt:    r.PublishedAt,
			}
		}
		c.JSON(http.StatusOK, result)
	}
}

func ApiGitRepositories(cfg *config.Config) gin.HandlerFunc {
	client := forgejo.NewClient(cfg.ForgejoURL)

	return func(c *gin.Context) {
		repos, err := client.GetPublicGitRepositories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		result := make([]dto.GitRepositoryResponse, len(repos))
		for i, repo := range repos {
			parts := strings.SplitN(repo.FullName, "/", 2)
			if len(parts) != 2 {
				continue
			}

			commits, err := client.GetCommits(parts[0], parts[1], 3)
			if err != nil {
				commits = []forgejo.GitCommit{}
			}

			commitDTOs := make([]dto.GitCommitResponse, len(commits))
			for j, c := range commits {
				commitDTOs[j] = dto.GitCommitResponse{
					URL:     c.HTMLURL,
					SHA:     c.SHA[:7],
					Message: c.Commit.Message,
					Author:  c.Commit.Author.Name,
					Date:    c.Commit.Author.Date.Format("2006-01-02"),
				}
			}

			result[i] = dto.GitRepositoryResponse{
				Name:    repo.FullName,
				URL:     repo.HTMLURL,
				Commits: commitDTOs,
			}
		}
		c.JSON(http.StatusOK, result)
	}
}
