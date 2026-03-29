SHELL := /bin/bash

.PHONY: start db-up db-seed backend frontend

start:
	./scripts/start.sh

db-up:
	docker compose up -d postgres

db-seed:
	./scripts/bootstrap-db.sh

backend:
	cd backend && API_ADDRESS=:8080 DATABASE_URL=postgres://nexus:nexus@localhost:5432/nexus?sslmode=disable FRONTEND_URL=http://localhost:3000 GOMODCACHE=/tmp/gomodcache GOCACHE=/tmp/gocache go run ./cmd/api

frontend:
	cd frontend && API_BASE_URL=http://localhost:8080 npm run dev -- --host 0.0.0.0 --port 3000
