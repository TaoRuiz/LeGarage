.PHONY: dev prod down build

dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build
	cd web-portal/frontend && npm run dev

prod:
	docker compose up -d --build

down:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml down

build:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml build