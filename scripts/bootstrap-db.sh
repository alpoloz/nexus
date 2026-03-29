#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT_DIR}"

POSTGRES_SERVICE="${POSTGRES_SERVICE:-postgres}"
POSTGRES_DB="${POSTGRES_DB:-nexus}"
POSTGRES_USER="${POSTGRES_USER:-nexus}"

echo "Starting Postgres container..."
if ! docker info >/dev/null 2>&1; then
  echo "Docker daemon is not running. Start Docker Desktop (or another Docker daemon) and retry." >&2
  exit 1
fi

docker compose up -d "${POSTGRES_SERVICE}"

echo "Waiting for Postgres to accept connections..."
until docker compose exec -T "${POSTGRES_SERVICE}" pg_isready -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" >/dev/null 2>&1; do
  sleep 1
done

echo "Applying migrations and seed data..."
for migration in "${ROOT_DIR}"/backend/migrations/*.sql; do
  echo "  -> $(basename "${migration}")"
  docker compose exec -T "${POSTGRES_SERVICE}" psql -v ON_ERROR_STOP=1 -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" < "${migration}" >/dev/null
done

echo "Database is ready."
