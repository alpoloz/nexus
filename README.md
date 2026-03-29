# Nexus

Nexus is a personal career portal with two read-only surfaces:

- `Engineering Blogs`: imported engineering blog posts stored in Postgres and linked to their external source URLs
- `Startups`: imported startup records stored in Postgres and presented as a searchable, filterable list

The MVP is intentionally read-only. There are no users, auth flows, or saved-state features yet, but the schema and service boundaries are designed so those can be added later without reshaping the core read models.

## Workspace layout

- `design/`: visual references and design system notes
- `frontend/`: Remix + React + TypeScript application
- `backend/`: Go + Echo + GORM API service
- `docs/`: architecture and API blueprint

## Local workflow

Run the full stack with:

```sh
make start
```

That command will:

1. start Postgres with Docker Compose
2. apply all SQL migrations in `backend/migrations/`
3. seed the database with the sample blog and startup records
4. start the Echo API on `http://localhost:8080`
5. start the Remix app on `http://localhost:3000`

Prerequisite: Docker must be running so Postgres can be started through `docker compose`.

Useful individual targets:

- `make db-up`
- `make db-seed`
- `make backend`
- `make frontend`

## Current scaffold

This repository now includes:

- a project structure for frontend and backend
- initial Postgres schema and seed data
- backend route and repository skeletons for read-only APIs
- Remix routes for `/blogs` and `/startups`
- shared project docs for architecture and API behavior

Frontend dependencies were installed during scaffolding. Backend dependencies are resolved through Go modules at runtime.
