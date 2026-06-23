FROM registry.access.redhat.com/ubi9/go-toolset:1.25 AS builder
USER 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY clowder-e2e.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /clowder-e2e

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.8
WORKDIR /
COPY --from=builder /clowder-e2e .
USER 65534:65534
CMD ["./clowder-e2e"]
