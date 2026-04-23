FROM golang:1.26-alpine AS builder

WORKDIR /src

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache ca-certificates tzdata

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -trimpath -ldflags="-s -w" -o /out/k8s-stats ./cmd/app

FROM alpine:3.23.4

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata curl && \
    addgroup -S app && adduser -S -G app app

COPY --from=builder /out/k8s-stats /app/k8s-stats

USER app

HEALTHCHECK --interval=15s --timeout=3s --start-period=20s --retries=3 \
  CMD curl -f http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["/app/k8s-stats"]
