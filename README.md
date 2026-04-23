# k8s-pobeda

Kubernetes stats backend with local Docker Compose run, Helm chart deployment, and GitHub Actions image build to GHCR.

## What is included

- Backend API in `backend/`
- OpenAPI spec in `backend/openapi.yml`
- Local run via `docker-compose.yml`
- Helm chart in `helm/k8s-pobeda-backend/`
- GitHub CI image workflow in `.github/workflows/backend-image.yml`

## Prerequisites

- Docker + Docker Compose
- Kubernetes cluster access (`~/.kube/config`) for local run
- Helm 3.x for chart deployment

## Run locally with Docker Compose

From repository root:

```bash
docker compose up --build -d
```

Check API:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/api/v1/stats
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

Workflow file: `.github/workflows/backend-image.yml`.

- On push to `main` and tags `v*`, CI builds and pushes image to GHCR.
- On pull requests, CI builds image without push.

Published image format:

```text
ghcr.io/<github-owner>/k8s-pobeda-backend:<tag>
```

## Backend configuration

Main env vars:

- `HTTP_ADDR` (default `:8080`)
- `CORS_ALLOW_ORIGIN` (default `*`)

See `backend/.env.example` and backend details in `backend/README.md`.
