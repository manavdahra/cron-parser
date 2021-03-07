# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BUILD_DIR="${ROOT_DIR}/bin"
BUILD_DIR_SHARE="${BUILD_DIR}/share"
BINARY=deliveroo-cron
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux
ARCHITECTURES=386 amd64

default: build

all: test clean build

test:
	go test ./... -count=1

build:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${BINARY}

build_all:
	mkdir -p ${BUILD_DIR}
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o ${BUILD_DIR_SHARE}/$(BINARY)-$(GOOS)-$(GOARCH))))

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

.PHONY: check test clean build build_all all