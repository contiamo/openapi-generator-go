default: container-build

all: bin/openapi-generator-go

bin/openapi-generator-go: $(shell find ./ -name "*.go")
	go build -o $@ ./cmd/openapi-generator-go

install: bin/openapi-generator-go
	cp $< $(shell go env GOPATH)/bin/

CONTAINER_ENGINE=podman
container-build:
	$(CONTAINER_ENGINE) run --rm \
	-v $(shell pwd):/app \
	-w /app \
	-v go-build-cache:/root/.cache/go-build \
	-v go-mod-cache:/go/pkg/mod \
	golang:1.15 make all

