FROM golang:1.12-alpine AS golang-base

FROM golang-base AS builder
WORKDIR /go/src/github.com/issho-ni/issho
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN apk add bash build-base curl git
COPY go.mod go.sum ./
RUN go mod download

FROM builder AS installer
COPY api api/
COPY cmd cmd/
COPY internal internal/
COPY mock mock/
RUN go test ./... \
    && go install ./cmd/...

FROM golang-base AS certs
ARG CERT_HOST=localhost
RUN go run /usr/local/go/src/crypto/tls/generate_cert.go --host ${CERT_HOST} \
    && mv cert.pem key.pem /

FROM scratch AS base
ENV TLS_KEY=/key.pem
ENV TLS_CERT=/cert.pem
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=certs /cert.pem /key.pem /

FROM base
ARG COMMAND
COPY --from=installer /go/bin/${COMMAND} /${COMMAND}
ENTRYPOINT [ "/${COMMAND}" ]
