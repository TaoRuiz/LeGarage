# Le Garage
> Work in progress.

## Stack

| Service              | Role                        |
|----------------------|-----------------------------|
| Caddy                | Reverse proxy + TLS         |
| Forgejo              | Git + CI/CD                 |
| Authentik            | SSO (OIDC)                  |
| Prometheus + Grafana | Metrics                     |
| Portal               | Landing + changelog service |

## Setup

```bash
# Production
cp .env.example .env

# Development
cp .env.dev.example .env
```

Generate a secret key for Authentik:

```bash
# Required for AUTHENTIK_SECRET_KEY in .env file
openssl rand -base64 36
```

## Run

```bash
# Production
docker compose up -d

# Development (plain HTTP, services exposed on localhost ports)
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

For local development, services are reachable directly via plain HTTP (no TLS, no domain needed):

```
http://localhost:3000  -> Forgejo
http://localhost:9000  -> Authentik
http://localhost:3001  -> Grafana
http://localhost:8080  -> Portal
```

```
# Visit to create admin account
http://localhost:9000/if/flow/initial-setup/

# Then visit
http://localhost:9000/if/admin/
```