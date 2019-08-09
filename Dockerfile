FROM golang:1.12-alpine AS golang-base

FROM golang-base AS builder
WORKDIR $GOPATH/src/github.com/issho-ni/issho
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN apk --update --no-cache add build-base git
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install ./cmd/...

FROM golang-base AS certs
ARG CERT_HOST=localhost
RUN go run /usr/local/go/src/crypto/tls/generate_cert.go --host $CERT_HOST \
    && mv cert.pem key.pem /

FROM scratch AS base
ENTRYPOINT ["/server"]
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=certs /cert.pem /key.pem /
ENV TLS_CERT=/cert.pem TLS_KEY=/key.pem

FROM base AS graphql
COPY --from=builder /go/bin/graphql /server

FROM base AS ninka
COPY --from=builder /go/bin/ninka /server

FROM base AS ninshou
COPY --from=builder /go/bin/ninshou /server

FROM base AS shinninjou
COPY --from=builder /go/bin/shinninjou /server

FROM base AS youji
COPY --from=builder /go/bin/youji /server

FROM scratch AS cache
COPY --from=graphql /cert.pem /
COPY --from=ninka /cert.pem /
COPY --from=ninshou /cert.pem /
COPY --from=shinninjou /cert.pem /
COPY --from=youji /cert.pem /
