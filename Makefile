# Golang Makefile for package golang-stress-tester

VERSION ?= v0.0.1
PACKAGE_NAME ?= $(shell grep "^module " go.mod | sed 's/^module\s//')
RUN_ARGS ?= -h
BUILD_MAIN_FILE = main.go
GO_CMD = go
GOAUTODOC_CMD = $(GO_CMD) run main.go
ECHO = "/usr/bin/echo"
# GOAUTODOC_CMD = $(shell go env GOPATH)/bin/goautodoc
TIME = $(shell date)
LDFLAGS = -X '$(PACKAGE_NAME)/config.PackageName=$(PACKAGE_NAME)' -X '$(PACKAGE_NAME)/config.Version=$(VERSION)'

build-linux-amd64:
	echo "Compiling linux-amd64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/amd64/$(PACKAGE_NAME)-linux-amd64 $(BUILD_MAIN_FILE)

build-linux-arm64:
	echo "Compiling linux-arm64"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/arm64/$(PACKAGE_NAME)-linux-arm64 $(BUILD_MAIN_FILE)

build-linux-arm:
	echo "Compiling linux-arm"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/arm/$(PACKAGE_NAME)-linux-arm $(BUILD_MAIN_FILE)

build-windows-amd64:
	echo "Compiling linux-arm64"
	CGO_ENABLED=0 GOOS=windows $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/windows/amd64/$(PACKAGE_NAME)-windows-amd64 $(BUILD_MAIN_FILE)

run:
	echo "Running the package..."
	$(GO_CMD) run -ldflags="$(LDFLAGS)" $(BUILD_MAIN_FILE) $(RUN_ARGS)

build-docs:
	$(GOAUTODOC_CMD) -config ./.autodoc/config.root.yml > ./README.md
	$(GOAUTODOC_CMD) -package ./config > ./config/README.md
	$(GOAUTODOC_CMD) -package ./doc > ./doc/README.md
	$(GOAUTODOC_CMD) -package ./templates > ./templates/README.md
	$(GOAUTODOC_CMD) -package ./templates/md > ./templates/md/README.md

show-info:
	$(ECHO) "Package:    '${PACKAGE_NAME}'" > /dev/null
	$(ECHO) "Version:    '${VERSION}'" > /dev/null
	$(ECHO) "Build file: '${BUILD_MAIN_FILE}'" > /dev/null
	$(ECHO) "Go command: '${GO_CMD}'" > /dev/null
	$(ECHO) "Time:       '$(TIME)'" > /dev/null

build-all: build-linux-amd64 build-linux-arm64 build-linux-arm build-windows-amd64
