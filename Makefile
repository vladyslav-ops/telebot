APP=$(shell basename $(shell git remote get-url origin 2>/dev/null || echo "telebot") .git)
# GitHub Container Registry (ghcr.io) — used by the CI/CD pipeline.
REGISTRY?=ghcr.io/vladyslav-ops
# Version = latest git tag + short commit sha, e.g. v1.0.0-106879e
VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v1.0.0")-$(shell git rev-parse --short HEAD)

# Host platform / architecture (the machine running `make`).
HOSTOS=$(shell go env GOOS)
HOSTARCH=$(shell go env GOARCH)

# Build target defaults to the host; per-platform targets override these.
TARGETOS?=$(HOSTOS)
TARGETARCH?=$(HOSTARCH)

IMAGE_TAG=$(REGISTRY)/$(APP):$(VERSION)-$(TARGETOS)-$(TARGETARCH)

format:
	gofmt -s -w ./

lint:
	golint ./...

test:
	go test -v ./...

get:
	go get ./...

build: format get
	CGO_ENABLED=0 GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build \
		-v -o telebot -ldflags "-X=github.com/vladyslav-ops/telebot/cmd.appVersion=$(VERSION)"

## Cross-platform build targets (Linux, ARM, macOS, Windows)
linux:
	$(MAKE) build TARGETOS=linux TARGETARCH=amd64

arm:
	$(MAKE) build TARGETOS=linux TARGETARCH=arm64

macos:
	$(MAKE) build TARGETOS=darwin TARGETARCH=arm64

windows:
	$(MAKE) build TARGETOS=windows TARGETARCH=amd64

## image: build the container image for the HOST platform/architecture.
## Standard `docker build` (BuildKit) — no buildx; cross-compiles via build-args.
image:
	docker build . -t $(IMAGE_TAG) \
		--build-arg TARGETOS=$(TARGETOS) \
		--build-arg TARGETARCH=$(TARGETARCH)

push:
	docker push $(IMAGE_TAG)

clean:
	rm -rf telebot
	docker rmi $(IMAGE_TAG)

.PHONY: format lint test get build linux arm macos windows image push clean
