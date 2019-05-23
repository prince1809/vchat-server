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

install:
	@echo You must be root to install

start-docker: ## Starts the docker containers for local development.
ifeq ($(IS_CI),false)
	@echo Starting docker containers

	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/mattermost-mysql$$ | wc -l) -eq 0 ]; then \
		echo starting mattermost-mysql; \
		docker run --name mattermost-mysql -p 3306:3306 \
			-e MYSQL_ROOT_PASSWORD=mostest \
			-e MYSQL_USER=mmuser \
			-e MYSQL_USER_PASSWORD=mostest \
			-e MYSQL_DATABASE=mattermost_test \
			-d mysql:5.7 > /dev/null; \
	elif [ $(shell docker ps --no-trunc --quiet --filter name=^/mattermost-mysql$$ | wc -l) -eq 0 ]; then \
		echo restarting mattermost-postgres; \
		docker start mattermost-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/mattermost-postgres$$ | wc -l) -eq 0 ]; then \
		echo starting mattermost-postgres; \
		docker run --name mattermost-postgres -p 5432:5432 \
			-e POSTGRES_USER=mmuser \
			-e POSTGRES_PASSWORD=mostest \
			-e POSTGRES_DB=mattermost_test \
			-d postgres:9.4 > /dev/null; \
		elif [ $(shell docker ps --no-trunc --quiet --filter name=^/mattermost-postgres$$ | wc -l) -eq 0 ]; then \
			echo restarting mattermost-postgres; \
			docker start mattermost-postgres > /dev/null; \
		fi

	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/mattermost-inbucket$$ | wc -l) -eq 0 ]; then \
		echo starting mattermost-inbucket; \
		docker run --name mattermost-inbucket -p 9000:10080 -p 2500:10025 -d jhillyerd/inbucket:release-1.2.0 > /dev/null; \
	elif [ $(shell docker ps --no-trunc --quiet --filter name=^/mattermost-inbucket$$ | wc -l) -eq 0 ]; then \
		echo restarting mattermost-inbucket; \
		docker start mattermost-inbucket > /dev/null; \
	fi
endif

run-server: start-docker ## Starts the server.
	@echo Running mattermost for development

	mkdir -p $(BUILD_WEBAPP_DIR)/dist/files
	$(GO) run $(GOFLAGS) -ldflags '$(LDFLAGS)' $(PLATFORM_FILES)
