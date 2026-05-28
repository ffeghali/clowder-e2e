# clowder-e2e

A minimal Go service used for end-to-end testing of [Clowder](https://github.com/RedHatInsights/clowder), the Red Hat Insights operator that manages cloud.redhat.com application deployments on OpenShift/Kubernetes.

## What it does

clowder-e2e is a lightweight HTTP server that validates Clowder's runtime configuration injection. It reads its configuration from `app-common-go` (the standard Clowder client library) and starts two servers:

- **Application server** on the Clowder-configured public port — serves a simple handler at `/`, `/healthz`, and `/api/puptoo/`
- **Metrics server** on the Clowder-configured metrics port — exposes Prometheus metrics

If a command-line argument is passed, it prints a message and exits instead of starting the servers. This allows testing both the service startup path and the CLI argument path in e2e scenarios.

## Prerequisites

- Go 1.19+
- A Clowder-managed environment (provides the `cdappconfig.json` configuration file)

## Building

```bash
go build -o clowder-e2e
```

### Container image

The project includes a multi-stage Dockerfile based on UBI9:

```bash
docker build -t clowder-e2e .
```

## Running

The binary expects Clowder to inject its configuration via the `ACG_CONFIG` environment variable pointing to a `cdappconfig.json` file:

```bash
ACG_CONFIG=/path/to/cdappconfig.json ./clowder-e2e
```

## Project structure

```
.
├── clowder-e2e.go   # Application entrypoint
├── Dockerfile       # Multi-stage container build (UBI9)
├── go.mod           # Go module definition
└── go.sum           # Dependency checksums
```
