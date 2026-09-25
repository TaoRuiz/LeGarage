package webhook

import (
	"log"
	"net/http"
	"time"

	"github.com/TaoRuiz/LeGarage/web-portal/internal/store"
	"github.com/gin-gonic/gin"
)

type gitReleasePayload struct {
	Action  string `json:"action"`
	Release struct {
		TagName     string    `json:"tag_name"`
		Name        string    `json:"name"`
		HTMLURL     string    `json:"html_url"`
		PublishedAt time.Time `json:"published_at"`
		Author      struct {
			Login string `json:"login"`
		} `json:"author"`
	} `json:"release"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func HandleGitRelease(s *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload gitReleasePayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}

		if payload.Action != "published" {
			c.JSON(http.StatusOK, gin.H{"status": "ignored"})
			return
		}

		r := store.Release{
			Repo:        payload.Repository.FullName,
			TagName:     payload.Release.TagName,
			Name:        payload.Release.Name,
			Author:      payload.Release.Author.Login,
			HTMLURL:     payload.Release.HTMLURL,
			PublishedAt: payload.Release.PublishedAt,
			ReceivedAt:  time.Now(),
		}

		if err := s.InsertRelease(r); err != nil {
			log.Printf("store error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "store error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "recorded"})
	}
}
