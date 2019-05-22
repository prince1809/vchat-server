.PHONY: run-server start-docker

ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

IS_CI ?= false
# Build Flags
BUILD_NUMBER ?= $(BUILD_NUMBER:)
BUILD_DATE = $(shell date -u)
BUILD_HASH = $(shell git rev-parse HEAD)
# If we don't set the build number it defaults to dev
ifeq ($(BUILD_NUMBER),)
    BUILD_NUMBER := dev
endif
BUILD_WEBAPP_DIR ?= ../vchat-webapp

# Golang Flags
GOPATH ?= $(shell go env GOPATH)
GOFLAGS ?= $(GOFLAGS:)
GO=go
DELVE=dlv

PLATFORM_FILES="./cmd/mattermost/main.go"

# Output paths
DIST_ROOT=dist
DIST_PATH=$(DIST_ROOT)/mattermost

# Tests
TESTS=.

start-docker: ## Starts the docker containers for local development.
ifeq ($(IS_CI),false)
	@echo Starting docker containers
endif

run-server: start-docker ## Starts the server.
	@echo Running mattermost for development

	mkdir -p $(BUILD_WEBAPP_DIR)/dist/files
	$(GO) run $(GOFLAGS) -ldflags '$(LDFLAGS)' $(PLATFORM_FILES)
