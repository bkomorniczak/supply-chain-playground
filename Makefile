SHELL := /bin/sh

APP_DIR := app
BIN := server
CMD := ./cmd/server
IMAGE ?= ci-supplychain-playground
TAG ?= dev

COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)
BUILD_TIME ?= $(shell date -u + "%Y-%m-%dT%H:%M:%SZ")

.PHONY: help tidy test build docker-build docker-run clean docker-smoke

help:
	@echo "Targets:"
	@echo "	tidy		    - go mod tidy in app"
	@echo "test				- run unit tests"
	@echo "  build        - build local binary (app/server)"
	@echo "  run          - run locally (PORT=8080 by default)"
	@echo "  docker-build - build container image with metadata"
	@echo "  docker-run   - run container locally"
	@echo "  docker-smoke - quick curl checks against container"
	@echo "  clean        - remove local build artifacts"


tidy:
	cd $(APP_DIR) && go mod tidy

test:
	cd $(APP_DIR) && go test ./...

build:
	cd $(APP_DIR) && go build -o $(BIN) $(CMD)

run: 
	cd $(APP_DIR) && PORT=8080 ./$(BIN)

docker-build:
	docker build \
		-f docker/Dockerfile \
		--build-arg COMMIT="$(COMMIT)" \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		-t $(IMAGE):$(TAG) \
		.

docker-run:
	docker run --rm -p 8080:8080 \
		-e PORT=8080 \
		$(IMAGE):$(TAG)

docker-smoke:
	@echo "Health:"
	@curl -fsS http://localhost:8080/healthz | sed 's/.*/  &/' || true
	@echo "Version:"
	@curl -fsS http://localhost:8080/version | sed 's/.*/  &/' || true
	@echo "Echo:"
	@curl -fsS "http://localhost:8080/echo?msg=hello" | sed 's/.*/  &/' || true

clean:
	rm -f $(APP_DIR)/$(BIN)