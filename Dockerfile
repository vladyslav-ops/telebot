FROM golang:1.21 AS builder

WORKDIR /go/src/app

COPY . .

ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN go get ./... && \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -v -o telebot -ldflags "-X=github.com/your-username/telebot/cmd.appVersion=v1.0.0"

FROM scratch

WORKDIR /

COPY --from=builder /go/src/app/telebot .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./telebot", "start"]