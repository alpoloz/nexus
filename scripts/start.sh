#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT_DIR}"

API_ADDRESS="${API_ADDRESS:-:8080}"
API_BASE_URL="${API_BASE_URL:-http://localhost:8080}"
FRONTEND_URL="${FRONTEND_URL:-http://localhost:3000}"
DATABASE_URL="${DATABASE_URL:-postgres://nexus:nexus@localhost:5432/nexus?sslmode=disable}"

"${ROOT_DIR}/scripts/bootstrap-db.sh"

cleanup() {
  local exit_code=$?

  if [[ -n "${BACKEND_PID:-}" ]] && kill -0 "${BACKEND_PID}" >/dev/null 2>&1; then
    kill "${BACKEND_PID}" >/dev/null 2>&1 || true
  fi

  if [[ -n "${FRONTEND_PID:-}" ]] && kill -0 "${FRONTEND_PID}" >/dev/null 2>&1; then
    kill "${FRONTEND_PID}" >/dev/null 2>&1 || true
  fi

  wait "${BACKEND_PID:-}" >/dev/null 2>&1 || true
  wait "${FRONTEND_PID:-}" >/dev/null 2>&1 || true

  exit "${exit_code}"
}

trap cleanup INT TERM EXIT

echo "Starting backend on ${API_ADDRESS}..."
(
  cd "${ROOT_DIR}/backend"
  export API_ADDRESS DATABASE_URL FRONTEND_URL GOMODCACHE=/tmp/gomodcache GOCACHE=/tmp/gocache
  exec go run ./cmd/api
) &
BACKEND_PID=$!

echo "Starting frontend on ${FRONTEND_URL}..."
(
  cd "${ROOT_DIR}/frontend"
  export API_BASE_URL
  exec npm run dev -- --host 0.0.0.0 --port 3000
) &
FRONTEND_PID=$!

echo "Nexus is starting."
echo "  Frontend: ${FRONTEND_URL}"
echo "  Backend:  ${API_BASE_URL}"

while true; do
  if ! kill -0 "${BACKEND_PID}" >/dev/null 2>&1; then
    wait "${BACKEND_PID}"
    exit $?
  fi

  if ! kill -0 "${FRONTEND_PID}" >/dev/null 2>&1; then
    wait "${FRONTEND_PID}"
    exit $?
  fi

  sleep 2
done
