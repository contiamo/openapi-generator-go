ARCH=amd64
OS=linux
CONTAINER_ENGINE=docker
.GIT_COMMIT=$(shell git rev-parse HEAD)
.GIT_VERSION=$(shell git describe --tags --always --dirty 2>/dev/null)
.GIT_UNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(.GIT_UNTRACKEDCHANGES),)
	.GIT_VERSION := $(.GIT_VERSION)-$(shell date +"%s")
endif

default: container-build

binary: bin/openapi-generator-go

bin/openapi-generator-go: $(shell find ./ -name "*.go")
	go build -v -o $@ -tags 'osusergo netgo' -v -ldflags "\
		-X github.com/contiamo/openapi-generator-go/cmd/openapi-generator-go/cmd.GitCommit=${.GIT_COMMIT} \
		-X github.com/contiamo/openapi-generator-go/cmd/openapi-generator-go/cmd.Version=${.GIT_VERSION}" \
	./cmd/openapi-generator-go

install: bin/openapi-generator-go
	cp $< $(shell go env GOPATH)/bin/

container-build:
	$(CONTAINER_ENGINE) run --rm \
	-v $(shell pwd):/app \
	-e GOOS=$(OS) \
	-e GOARCH=$(ARCH) \
	-w /app \
	-v go-build-cache:/root/.cache/go-build \
	-v go-mod-cache:/go/pkg/mod \
	golang:1.15 make binary

crosscompile:
	mkdir -p bin/linux/{arm,arm64,amd64} bin/darwin/amd64
	$(MAKE) container-build OS=linux ARCH=amd64 && mv bin/openapi-generator-go bin/linux/amd64/
	$(MAKE) container-build OS=linux ARCH=arm && mv bin/openapi-generator-go bin/linux/arm/
	$(MAKE) container-build OS=linux ARCH=arm64 && mv bin/openapi-generator-go bin/linux/arm64/
	$(MAKE) container-build OS=darwin ARCH=amd64 && mv bin/openapi-generator-go bin/darwin/amd64/
