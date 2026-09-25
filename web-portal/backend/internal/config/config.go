package config

import "os"

type Config struct {
	WebPortalFrontURL string
	ForgejoURL        string
	ForgejoToken      string
	AuthentikURL      string
	GrafanaURL        string
	PrometheusURL     string
}

func Load() *Config {
	return &Config{
		WebPortalFrontURL: getEnv("WEBPORTALFRONT_URL", "http://localhost:5173"),
		ForgejoURL:        getEnv("FORGEJO_URL", "http://localhost:3000"),
		AuthentikURL:      getEnv("AUTHENTIK_URL", "http://localhost:9000"),
		GrafanaURL:        getEnv("GRAFANA_URL", "http://localhost:3001"),
		PrometheusURL:     getEnv("PROMETHEUS_URL", "http://localhost:9090"),
		ForgejoToken:      getEnv("FORGEJO_TOKEN", ""),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
