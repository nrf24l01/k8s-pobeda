# Kubernetes Stats Backend

Small Go API that reads Kubernetes cluster stats and exposes them over HTTP.

## Endpoints

- `GET /healthz` - liveness check
- `GET /api/v1/stats` - cluster statistics snapshot

OpenAPI spec: `openapi.yml`

## What stats are returned

- time from latest pending pod
- time from latest pending node
- node count
- pod count
- namespace count
- allocated resources
- allocatable resources

Response intentionally contains only numeric fields for stats values.
Time values are Unix epoch timestamps (`*Epoch`) and derived age in seconds (`*AgeSeconds`).

## Prerequisites

- Go `1.26+`
- Access to a Kubernetes cluster:
  - in-cluster config when running inside Kubernetes, or
  - valid local kubeconfig (default kubeconfig loading rules)

## Run locally

```bash
go mod tidy
go run ./cmd/app
```

By default, server listens on `:8080`.
The app auto-loads variables from `.env` if the file exists.

Set custom bind address:

```bash
HTTP_ADDR=":8080" go run ./cmd/app
```

Set CORS allowed origin (default `*`):

```bash
CORS_ALLOW_ORIGIN="http://localhost:3000" go run ./cmd/app
```

Environment examples are available in `.env.example`.

Test endpoints:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/api/v1/stats
```

## Build binary

```bash
go build -o bin/k8s-stats ./cmd/app
./bin/k8s-stats
```

## Run with Docker

Build image:

```bash
docker build -t k8s-pobeda-backend:latest .
```

Run container:

```bash
docker run --rm -p 8080:8080 -e HTTP_ADDR=":8080" k8s-pobeda-backend:latest
```

If running outside Kubernetes, mount kubeconfig:

```bash
docker run --rm -p 8080:8080 \
  -e HTTP_ADDR=":8080" \
  -v "$HOME/.kube:/home/nonroot/.kube:ro" \
  -e KUBECONFIG=/home/nonroot/.kube/config \
  k8s-pobeda-backend:latest
```

## Verify

```bash
go test ./...
```
