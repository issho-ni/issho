FROM golang:1.12-alpine AS base

FROM base AS certs
ARG CERT_HOST=localhost
RUN go run /usr/local/go/src/crypto/tls/generate_cert.go --host $CERT_HOST && \
    mv cert.pem key.pem /

FROM base AS builder
WORKDIR $GOPATH/src/github.com/issho-ni/issho
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN apk --update --no-cache add build-base git
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install ./cmd/...

FROM scratch
ENTRYPOINT ["/server"]
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=certs /cert.pem /key.pem /
ENV TLS_CERT=/cert.pem TLS_KEY=/key.pem
ARG COMMAND
COPY --from=builder /go/bin/${COMMAND} /server
