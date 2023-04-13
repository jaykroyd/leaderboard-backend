include mocks.mk
include pg.mk
-include .env

SHELL := /bin/sh
PACKAGES := $(shell go list ./...)
ifdef DEBUG
$(info PACKAGES = $(PACKAGES))
$(info SOURCES = $(SOURCES))
endif

################################################################################
## Build make targets
################################################################################

.PHONY: build
build:
	go build .

.PHONY: run-api
run-api:
	build
	leaderboard start --type=api

.PHONY: dev-run-api
dev-run-api:
	@make build
	./leaderboard start --type=api

################################################################################
## Tests
################################################################################

TESTABLE_PACKAGES := $$(go list ./... | egrep -v 'constants|mocks|testing')

.PHONY: test-unit
test-unit:
	echo 'Starting unit tests...'
	go test -timeout 50000ms -tags=unit -coverprofile=coverage.out ./...

.PHONY: test-integration
test-integration:
	echo 'Starting integration tests...'
	@GO111MODULE=on go test ${TESTABLE_PACKAGES} -tags=integration -coverprofile integration.coverprofile
	go test -tags=integration -coverprofile=coverage.out ./...
	echo 'Removing test environment...'

################################################################################
## Linters and formatters
################################################################################

.PHONY: fix
fix:
	go mod tidy

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run
	go run golang.org/x/lint/golint $(PACKAGES)
