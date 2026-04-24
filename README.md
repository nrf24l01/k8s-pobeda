# k8s-pobeda

Kubernetes stats app with frontend + backend, local Docker Compose run, Helm chart deployment, and GitHub Actions image build to GHCR.

## What is included

- Backend API in `backend/`
- Frontend app in `frontend/`
- OpenAPI spec in `backend/openapi.yml`
- Local run via `docker-compose.yml`
- Helm chart in `helm/k8s-pobeda-backend/` (frontend + backend + gateway)
- GitHub CI image workflows in `.github/workflows/backend-image.yml` and `.github/workflows/frontend-image.yml`

## Prerequisites

- Docker + Docker Compose
- Kubernetes cluster access (`~/.kube/config`) for local run
- Helm 3.x for chart deployment

## Run locally with Docker Compose

From repository root:

```bash
docker compose up --build -d
```

Routing in local compose:

- `/` -> frontend
- `/api` -> backend

Check:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/api/v1/stats
```

Open frontend:

```bash
xdg-open http://localhost:8080
```

Stop:

```bash
docker compose down
```

## Deploy to Kubernetes with Helm

Install/upgrade release:

```bash
helm upgrade --install k8s-pobeda-backend ./helm/k8s-pobeda-backend \
  --namespace k8s-pobeda \
  --create-namespace
```

Override image/tag (example):

```bash
helm upgrade --install k8s-pobeda-backend ./helm/k8s-pobeda-backend \
  --namespace k8s-pobeda \
  --set image.repository=ghcr.io/<owner>/k8s-pobeda-backend \
  --set image.tag=<tag>
```

## GitHub CI image build

Workflow files:

- `.github/workflows/backend-image.yml`
- `.github/workflows/frontend-image.yml`

- On push to `main` and tags `v*`, CI builds and pushes backend/frontend images to GHCR.
- On pull requests, CI builds images without push.

Published image formats:

```text
ghcr.io/<github-owner>/k8s-pobeda-backend:<tag>
ghcr.io/<github-owner>/k8s-pobeda-frontend:<tag>
```

## Backend configuration

Main env vars:

- `HTTP_ADDR` (default `:8080`)
- `CORS_ALLOW_ORIGIN` (default `*`)

See `backend/.env.example` and backend details in `backend/README.md`.

## Frontend build API path

Frontend build is configured to use `/api` as backend base path:

- in Docker image build via `NUXT_PUBLIC_API_BASE_URL=/api`
- in runtime env via `NUXT_PUBLIC_API_BASE_URL=/api`

## Helm routing

Helm chart deploys 3 services/components:

- frontend (`*-frontend`)
- backend (`*`)
- nginx gateway (`*-gateway`)

Gateway routes:

- `/` -> frontend
- `/api` and `/api/*` -> backend
