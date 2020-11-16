.CONTAINER_ENGINE=docker
.GIT_COMMIT=$(shell git rev-parse HEAD)
.GIT_VERSION=$(shell git describe --tags --always --dirty 2>/dev/null)
.GIT_UNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(.GIT_UNTRACKEDCHANGES),)
	.GIT_VERSION := $(.GIT_VERSION)-$(shell date +"%s")
endif


default: container-build

all: bin/openapi-generator-go

bin/openapi-generator-go: $(shell find ./ -name "*.go")
	go build -o $@ -tags 'osusergo netgo' -v -ldflags "\
		-X github.com/contiamo/openapi-generator-go/cmd/openapi-generator-go/cmd.GitCommit=${.GIT_COMMIT} \
		-X github.com/contiamo/openapi-generator-go/cmd/openapi-generator-go/cmd.Version=${.GIT_VERSION}" \
	./cmd/openapi-generator-go

install: bin/openapi-generator-go
	cp $< $(shell go env GOPATH)/bin/

container-build:
	$(.CONTAINER_ENGINE) run --rm \
	-v $(shell pwd):/app \
	-w /app \
	-v go-build-cache:/root/.cache/go-build \
	-v go-mod-cache:/go/pkg/mod \
	golang:1.15 make all

