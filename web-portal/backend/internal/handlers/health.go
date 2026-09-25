package handlers

import (
	"net/http"
	"time"
)

type platformToolsService struct {
	Name string
	URL  string
}

type PlatformToolsServiceStatus struct {
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
}

var httpClient = &http.Client{
	Timeout: 2 * time.Second,
}

func checkService(url string) bool {
	resp, err := httpClient.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode < 500
}
