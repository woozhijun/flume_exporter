unexport GOBIN

GO    ?= go
GOFMT ?= $(GO) fmt
GOPATH := $(shell $(GO) env GOPATH)
GOOPTS       ?=
GOHOSTOS     ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH   ?= $(shell $(GO) env GOHOSTARCH)
GO_VERSION   ?= $(shell $(GO) version)

PROMU := $(GOPATH)/bin/promu
pkgs   = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
DOCKER_IMAGE_NAME       ?= flume_exporter
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
TAG 					:= $(shell echo `if [ "$(TRAVIS_BRANCH)" = "master" ] || [ "$(TRAVIS_BRANCH)" = "" ] ; then echo "latest"; else echo $(TRAVIS_BRANCH) ; fi`)

export GO111MODULE=on
export GOPROXY=https://goproxy.io

all: format test build

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

test:
	@echo ">> running tests"
	@$(GO) test -short $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build: promu
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)

crossbuild: promu
	@echo ">> crossbuilding binaries"
	@$(PROMU) crossbuild

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball $(BIN_DIR)

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

push:
	@echo ">> pushing docker image, $(DOCKER_USERNAME),$(DOCKER_IMAGE_NAME),$(TAG)"
#	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)
	@docker tag "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" "$(DOCKER_USERNAME)/$(DOCKER_IMAGE_NAME):$(TAG)"
	@docker push "$(DOCKER_USERNAME)/$(DOCKER_IMAGE_NAME):$(TAG)"

promu:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get -u github.com/prometheus/promu

release: promu github-release
	@echo ">> pushing binary to github with ghr"
	@$(PROMU) crossbuild tarballs
	@$(PROMU) release .tarballs

github-release:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get -u github.com/aktau/github-release

.PHONY: all vet docker promu
