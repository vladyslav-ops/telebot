APP=$(shell basename $(shell git remote get-url origin 2>/dev/null || echo "telebot") .git)
REGISTRY?=your-dockerhub-username
VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v1.0.0")
TARGETOS?=linux
TARGETARCH?=amd64

format:
	gofmt -s -w ./

lint:
	golint ./...

test:
	go test -v ./...

get:
	go get ./...

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
		-v -o telebot -ldflags "-X=github.com/your-username/telebot/cmd.appVersion=${VERSION}"

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

clean:
	rm -rf telebot
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH} 2>/dev/null || true

.PHONY: format lint test get build image push clean