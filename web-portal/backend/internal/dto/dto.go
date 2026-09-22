package dto

import "time"

type PlatformToolsServiceResponse struct {
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
}

type GitCommitResponse struct {
	URL     string `json:"url"`
	SHA     string `json:"name"`
	Message string `json:"message"`
	Author  string `json:"author"`
	Date    string `json:"date"`
}

type GitRepositoryResponse struct {
	Name    string              `json:"name"`
	URL     string              `json:"url"`
	Commits []GitCommitResponse `json:"commits"`
}

type GitReleaseResponse struct {
	ID             int64     `json:"id"`
	RepositoryName string    `json:"repositoryName"`
	Tag            string    `json:"tag"`
	TagName        string    `json:"tagName"`
	Author         string    `json:"author"`
	URL            string    `json:"url"`
	PublishedAt    time.Time `json:"publishedAt"`
}
