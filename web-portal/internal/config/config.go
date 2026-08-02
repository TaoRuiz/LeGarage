package config

import "os"

type Config struct {
	ForgejoURL    string
	AuthentikURL  string
	GrafanaURL    string
	PrometheusURL string
}

func Load() *Config {
	return &Config{
		ForgejoURL:    getEnv("FORGEJO_URL", "http://localhost:3000"),
		AuthentikURL:  getEnv("AUTHENTIK_URL", "http://localhost:9000"),
		GrafanaURL:    getEnv("GRAFANA_URL", "http://localhost:3001"),
		PrometheusURL: getEnv("PROMETHEUS_URL", "http://localhost:9090"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
