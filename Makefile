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


start-docker: ## Starts the docker containers for local development.
ifeq ($(IS_CI),false)
	@echo Starting docker containers
endif

run-server: start-docker ## Starts the server.
	@echo Running mattermost for development

	mkdir -p $(BUILD_WEBAPP_DIR)/dist/files

