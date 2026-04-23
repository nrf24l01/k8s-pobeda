# AGENTS.md

## Project overview

This repository is a standard Go module.

Primary goals for any change:
- keep the code simple and idiomatic
- prefer small packages with clear ownership
- avoid unnecessary abstractions
- preserve backward-compatible behavior unless the task explicitly requires otherwise

## Tech stack

- Language: Go
- Module system: Go modules
- Build/test tooling: `go` CLI
- Preferred formatting: `gofmt`
- Preferred linting: `go vet` and `golangci-lint` when configured

## Repository structure

Typical layout used in this repo:

```text
.
├── cmd/                  # executable entrypoints
│   └── app/              # main package for the main binary
├── internal/             # private application code; do not import from outside this module
│   ├── app/              # app-specific orchestration
│   ├── domain/           # core business logic and entities
│   ├── service/          # use cases / business services
│   ├── transport/        # HTTP, gRPC, CLI, messaging adapters
│   ├── store/            # DB/repository implementations
│   └── config/           # config loading and validation
├── pkg/                  # public reusable packages, only if truly intended for external use
├── api/                  # OpenAPI, protobuf, or contract files if present
├── configs/              # sample config files
├── scripts/              # helper scripts for dev/CI
├── testdata/             # fixtures used by tests
├── go.mod
├── go.sum
└── README.md