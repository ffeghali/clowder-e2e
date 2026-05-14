FROM golang as compiler
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY clowder-e2e.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /clowder-e2e

FROM registry.access.redhat.com/ubi8/ubi-minimal
RUN microdnf install tar curl
COPY --from=compiler /clowder-e2e .
CMD ["./clowder-e2e"]