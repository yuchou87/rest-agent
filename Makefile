.PHONY: all
all: tidy lint build

SHELL := /bin/bash
DIRS=$(shell ls)
GO=go
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# include the common makefile
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# ROOT_DIR: root directory of the code base
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/. && pwd -P))
endif
# OUTPUT_DIR: The directory where the build output is stored.
ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/bin
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --abbrev=0 --dirty --always --tags | sed 's/-/./g')
endif

# Check if the tree is dirty. default to dirty(maybe u should commit?)
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)


BUILDFILE = "./cmd/cli/main.go"
BUILDAPP = "$(OUTPUT_DIR)/rest-agent"



# ==============================================================================
# Targets

## build: Build binaries by default
.PHONY: build
build: 
	@echo "$(shell go version)"
	@echo "===========> Building binary $(BUILDAPP) *[Git Info]: $(VERSION)-$(GIT_COMMIT)"
	@export CGO_ENABLED=0 && go build -o $(BUILDAPP) -ldflags "-s -w -X main.version=dev -X main.commit=$$(git rev-parse --short HEAD) -X main.date=$$(date +%FT%TZ)" $(BUILDFILE)

## tidy: tidy go.mod
.PHONY: tidy
tidy:
	@$(GO) mod tidy

## fmt: Run go fmt against code.
.PHONY: fmt
fmt:
	@$(GO) fmt ./...

## vet: Run go vet against code.
.PHONY: vet
vet:
	@$(GO) vet ./...

## lint: Run go lint against code.
.PHONY: lint
lint:
	@golangci-lint run -v ./...

## style: Code style -> fmt,vet,lint
.PHONY: style
style: fmt vet lint

## test: Run unit test
.PHONY: test
test: 
	@echo "===========> Run unit test"
	@$(GO) test ./... 

## cover: Run unit test with coverage
.PHONY: cover
cover: test
	@$(GO) test -cover

## go.clean: Clean all builds
.PHONY: clean
clean:
	@echo "===========> Cleaning all builds OUTPUT_DIR($(OUTPUT_DIR))"
	@-rm -vrf $(OUTPUT_DIR)
	@echo "===========> End clean..."
