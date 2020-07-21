PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION = 1.0.0
COMMIT := $(shell git log -1 --format='%H')
current_dir = $(shell pwd)

# TODO: Update the ldflags with the app, client & server names
ldflags = -X github.com/TsukiCore/cosmos-sdk/version.Name=tsuki \
	-X github.com/TsukiCore/cosmos-sdk/version.ServerName=tsukid \
	-X github.com/TsukiCore/cosmos-sdk/version.ClientName=tsukicli \
	-X github.com/TsukiCore/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/TsukiCore/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/tsukid
		go install $(BUILD_FLAGS) ./cmd/tsukicli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify
