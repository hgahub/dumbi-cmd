.PHONY: build

GIT_COMMIT := $(shell git rev-parse --short HEAD)

build:
	go build -ldflags "-X main.gitCommit=$(GIT_COMMIT)" -o dumbi .