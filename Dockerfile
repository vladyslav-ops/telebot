FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.26 AS builder

WORKDIR /go/src/app

COPY . .

ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN go get ./... && \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -v -o telebot -ldflags "-X=github.com/vladyslav-ops/telebot/cmd.appVersion=v1.0.0"

FROM scratch

WORKDIR /

COPY --from=builder /go/src/app/telebot .

COPY --from=builder /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./telebot", "start"]
