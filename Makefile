.PHONY: help build
define PRINT_HELP_PYSCRIPT
import re, sys

print("Usage:  make COMMAND\n\nCommands:")

for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("  %-20s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

help:  ## Print this help message and exit
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)

# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /usr/bin/env bash

# Commit Hash for CI build.
GIT_COMMIT := $(shell git rev-parse --short HEAD)

build: ## Application build
	go build -ldflags "-X main.commit=$(GIT_COMMIT)" -o dumbi .

sonar: ## SonarQube
	go test ./... -json > test-report.out
	go test ./... -coverprofile=coverage.out
	scripts/sonar.sh

test: ## Test
	go test ./...