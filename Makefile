
include .bingo/Variables.mk

.PROJECT_ROOT=$(shell pwd)
.OWNER=contiamo
.NAME=openapi-generator-go
REGISTRY=


GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_VERSION=$(shell git describe --tags --always --dirty 2>/dev/null)
.GIT_UNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(.GIT_UNTRACKEDCHANGES),)
	GIT_VERSION := $(GIT_VERSION)-$(shell date +"%s")
endif


GOBINS=$(shell go env GOBIN)

default: install

.PHONY: version
version:
	@echo $(GIT_VERSION)

## Build env
setup: $(GORELEASER) $(GOLANGCI_LINT) ## install the required build environment binaries


################################
################################
.PHONY: lint
lint: $(GOLANGCI_LINT) ## Verifies `golangci-lint` passes
	@echo "+ $@"
	@$(GOLANGCI_LINT) run ./...

.PHONY: lint-fix
lint-fix: $(GOLANGCI_LINT) ## Verifies `golangci-lint` passes
	@echo "+ $@"
	@$(GOLANGCI_LINT) run --fix ./...

.PHONY: fmt
fmt:
	@echo "+ $@"
	@gofmt -s -l . | tee /dev/stderr

test:
	@go test -cover ./...
	@cd pkg/generators/models/testdata/cases/oneof_mapped_discriminator/expected && go test -cover ./...
	@cd pkg/generators/models/testdata/cases/validation/expected && go test -cover ./...


${GOBINS}/${.NAME}:  $(shell find . -name '*.go') go.*
	@echo "+ $@"
	go install -tags 'osusergo netgo' -v -ldflags "\
		-X github.com/contiamo/openapi-generator-go/cmd/openapi-generator-go/cmd.GitCommit=${GIT_COMMIT} \
		-X github.com/contiamo/openapi-generator-go/cmd/openapi-generator-go/cmd.Version=${GIT_VERSION}" \
		./cmd/openapi-generator-go/main.go

install: ${GOBINS}/${.NAME}

dist: $(GORELEASER) $(shell find . -name '*.go') go.*
	$(GORELEASER) build --skip-validate --rm-dist


image: dist
	docker buildx build \
		--load \
		--platform linux/amd64 \
		--tag $(REGISTRY)$(.OWNER)/$(.NAME):$(GIT_VERSION) \
		.

image-multi: dist
	docker buildx build \
		--push \
		--platform linux/amd64,linux/arm64 \
		--tag $(REGISTRY)$(.OWNER)/$(.NAME):$(GIT_VERSION) \
		.
