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
BUILD_ENTERPRISE_DIR ?= ../enterprise
BUILD_ENTERPRISE ?= true
BUILD_ENTERPRISE_READY = false
BUILD_TYPE_NAME = team
BUILD_HASH_ENTERPRISE = none
LDAP_DATA ?= test
ifneq ($(wildcard $(BUILD_ENTERPRISE_DIR)/.),)
	ifeq ($(BUILD_ENTERPRISE),true)
		BUILD_ENTERPRISE_READY = true
		BUILD_TYPE_NAME = enterprise
		BUILD_HASH_ENTERPRISE = $(shell cd $(BUILD_ENTERPRISE_DIR) && git rev-parse HEAD)
	else
		BUILD_ENTERPRISE_READY = false
		BUILD_TYPE_NAME = team
	endif
else
	BUILD_ENTERPRISE_READY = false
	BUILD_TYPE_NAME = team
endif
BUILD_WEBAPP_DIR ?= ../vchat-webapp
BUILD_CLIENT = false
BUILD_HASH_CLIENT = independent
ifneq ($(wildcard $(BUILD_WEBAPP_DIR)/.),)
	ifeq ($(BUILD_CLIENT), true)
		BUILD_CLIENT = true
		BUILD_HASH_CLIENT = $(shell cd $(BUILD_WEBAPP_DIR) && git rev-parse HEAD)
	else
		BUILD_CLIENT = false
	endif
else
	BUILD_CLIENT = false
endif

# Golang Flags
GOPATH ?= $(shell go env GOPATH)
GOFLAGS ?= $(GOFLAGS:)
GO=go
DELVE=dlv
LDFLAGS += -X "github.com/prince1809/mattermost-server/model.BuildNumber=$(BUILD_NUMBER)"
LDFLAGS += -X "github.com/prince1809/mattermost-server/model.BuildDate=$(BUILD_DATE)"
LDFLAGS += -X "github.com/prince1809/mattermost-server/model.BuildHash=$(BUILD_HASH)"
LDFLAGS += -X "github.com/prince1809/mattermost-server/model.BuildHashEnterprise=$(BUILD_HASH_ENTERPRISE)"
LDFLAGS += -X "github.com/prince1809/mattermost-server/model.BuildEnterpriseReady=$(BUILD_ENTERPRISE_READY)"

# GOOS/GOARCH of the build host, used to determine whether we're cross-compiling or not
BUILDER_GOOS_ARCH="$(shell $(GO) env GOOS)_$(shell $(GO) env GOARCH)"

PLATFORM_FILES="./cmd/mattermost/main.go"

# Output paths
DIST_ROOT=dist
DIST_PATH=$(DIST_ROOT)/mattermost

# Tests
TESTS=.

TESTFLAGS ?= -short
TESTFLAGSEE ?= -short

# Packages lists
TE_PACKAGES=$(shell go list ./...|grep -v plugin_tests)

# Plugins Packages
PLUGIN_PACKAGES=mattermost-plugin-zoom mattermost-plugin-jira

# Prepare the enterprise build if exists. The IGNORE stuff is a hack to get the Makefile to exeute the commands outside a target
ifeq ($(BUILD_ENTERPRISE_READY),true)
	IGNORE:=$(shell echo Enterprise build selected, preparing)
	IGNORE:=$(shell rm -f imports/imports.go)
	INGORE:=$(shell cp $(BUILD_ENTERPRISE_DIR)/imports/imports.go imports/)
	IGNORE:=$(shell rm -f enterprise)
	IGNORE:=$(shell ln -s $(BUILD_ENTERPRISE_DIR) enterprise)
else
	IGNORE:=$(shell rm -f imports/imports.go)
endif

EE_PACKAGES=$(shell go list ./enterprise/...)

ifeq ($(BUILD_ENTERPRISE_READY),true)
	ALL_PACKAGES=$(TE_PACKAGES) $(EE_PACKAGES)
else
	ALL_PACKAGES=$(TE_PACKAGES)
endif

all: run ## Alias for 'run'.

include build/*.mk


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

	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/mattermost-minio$$ | wc -l) -eq 0 ]; then \
		echo starting mattermost-minio; \
		docker run --name mattermost-minio -p 9001:9000 -e "MINIO_ACCESS_KEY=minioaccesskey" \
		-e "MINIO_SSE_MASTER_KEY=my-minio-key:6368616e676520746869732070617373776f726420746f206120736563726574" \
		-e "MINIO_SECRET_KEY=miniosecretkey" -d minio/minio:RELEASE.2019-04-23T23-50-36Z server /data > /dev/null; \
		docker exec -it mattermost-minio /bin/sh -c "mkdir -p /data/mattermost-test" > /dev/null; \
	elif [ $(shell docker ps --no-trunc --quiet --filter name=^/mattermost-minio$$ | wc -l) -eq 0 ]; then \
		echo restarting mattermost-minio; \
		docker start mattermost-minio > /dev/null; \
	fi

ifeq ($(BUILD_ENTERPRISE_READY),true)
	@echo Ldap test user test.one
	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/mattermost-openldap$$ | wc -l) -eq 0 ]; then \
		echo starting mattermost-openldap; \
		docker run --name mattermost-openldap -p 389:389 -p 636:636 \
			-e LDAP_TLS_VERIFY_CLIENT="never" \
			-e LDAP_ORGANISATION="Mattermost Test" \
			-e LDAP_DOMAIN="mm.test.com" \
			-e LDAP_ADMIN_PASSWORD="mostest" \
			-d osixia/openldap:1.2.2 > /dev/null;\
		sleep 10; \
		docker cp tests/test-data.ldif mattermost-openldap:/test-data.ldif;\
		docker cp tests/qa-data.ldif mattermost-openldap:/qa-data.ldif;\
		docker exec -ti mattermost-openldap bash -c 'ldapadd -x -D "cn=admin,dc=mm,dc=test,dc=com" -w mostest -f /$(LDAP_DATA)-data.ldif';\
	elif [ $(shell docker ps | grep -ci mattermost-openldap) -eq 0 ]; then \
		echo restarting mattermost-openldap; \
		docker start matermost-openldpa > /dev/null; \
		sleep 10; \
	fi
endif
else
	@echo CI Build: skipping docker start
endif

govet: ## Runs govet against all packages.
	@echo Running GOVET
	$(GO) get golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
	$(GO) vet $(GOFLAGS) $(ALL_PACKAGES) || exit 1
	$(GO) vet -vettool=$(GOPATH)/bin/shadow $(GOFLAGS) $(ALL_PACKAGES) || exit 1

gofmt: ## Runs gofmt against all packages
	@echo Running GOFMT

	@for package in $(TE_PACKAGES); do \
		echo "Checking " $$package; \
		files=$$(go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			gofmt_output=$$(); \
			if [ "$$gofmt_output" ]; then \
				echo "$$gofmt_output"; \
				echo "gofmt failure"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "gofmt success"; \

ifeq ($(BUILD_ENTERPRISE_READY), true)
	$(GOPATH)/bin/megacheck $(EE_PACKAGES) || exit 1
endif

test-data: start-docker ## Add test data to the local instance.
	$(GO) run $(GOFLAGS) -ldflags '$(LDFLAGS)' $(PLATFORM_FILES) sampledata -w 1

	@echo You may need to restart the Mattermost server before using the folli

run-server: start-docker ## Starts the server.
	@echo Running mattermost for development

	mkdir -p $(BUILD_WEBAPP_DIR)/dist/files
	$(GO) run $(GOFLAGS) -ldflags '$(LDFLAGS)' $(PLATFORM_FILES) --disableconfigwatch | \
	    $(GO) run $(GOFLAGS) -ldflags '$(LDFLAGS)' $(PLATFORM_FILES) logs --logrus &


run-client: ## Runs the webapp.
	@echo Running mattermost client for development

	ln -nfs $(BUILD_WEBAPP_DIR)/dist clean
	cd $(BUILD_WEBAPP_DIR) && $(MAKE) run
