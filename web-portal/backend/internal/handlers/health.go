package handlers

import (
	"net/http"
	"time"
)

type ServiceStatus struct {
	Name    string
	Healthy bool
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
