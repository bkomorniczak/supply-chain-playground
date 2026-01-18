# syntax=docker/dockerfile:1.7

# --- Build stage ---
# Pin Go toolchain version so builds don't drift silently.
FROM golang:1.22.6-bookworm AS build

WORKDIR /src

# Build metadata injected at build time.
ARG COMMIT=dev
ARG BUILD_TIME=unknown

# Faster + more reproducible module download.
ENV CGO_ENABLED=0 \
    GOFLAGS="-trimpath" \
    GOMODCACHE=/go/pkg/mod

# Copy go.mod first to leverage Docker layer cache.
COPY app/go.mod app/go.sum* ./app/
WORKDIR /src/app
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Now copy the rest of the source.
COPY app/ .

# Run tests in the build stage (optional but useful as a baseline).
RUN --mount=type=cache,target=/go/pkg/mod \
    go test ./...

# Build the binary.
RUN --mount=type=cache,target=/go/pkg/mod \
    go build -ldflags "-s -w \
      -X ci-supplychain-playground/app/internal/version.commit=${COMMIT} \
      -X ci-supplychain-playground/app/internal/version.buildTime=${BUILD_TIME}" \
    -o /out/server ./cmd/server


# --- Runtime stage ---
# Distroless = minimal, no shell, nonroot by default.
FROM gcr.io/distroless/static-debian12:nonroot AS runtime

WORKDIR /app

# Copy the statically-linked binary.
COPY --from=build /out/server /app/server

# Expose the app port (informational; runtime still uses PORT env).
EXPOSE 8080

# Distroless runs as nonroot user already.
ENTRYPOINT ["/app/server"]