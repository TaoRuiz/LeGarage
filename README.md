## Setup

```bash
cp .env.dev.example .env
```

Edit `.env` — change the passwords, generate an Authentik secret key:

```bash
openssl rand -base64 36
```

## Run

```bash
# Development
make dev

# Production
make prod
```

### Local services (dev)
http://localhost:3000 -> Forgejo  
http://localhost:9000 -> Authentik  
http://localhost:3001 -> Grafana  
http://localhost:5173 -> Portal  
http://localhost:8080 -> Portal API  

### First run
Authentik admin setup
http://localhost:9000/if/flow/initial-setup/
