# issho-ni

[![Build Status](https://ci.issho-ni.co/api/v1/teams/main/pipelines/issho/badge)](https://ci.issho-ni.co/main/pipelines/issho)
[![Go Report Card](https://goreportcard.com/badge/github.com/issho-ni/issho)](https://goreportcard.com/report/github.com/issho-ni/issho)

## Requirements

- [Go][] 1.12
- [MongoDB][]
- [mkcert][]

## Development

### Setup

```bash
export GO111MODULE=on
mkcert localhost 127.0.0.1 ::1
```

### Environment Variables

#### `GRAPHQL_PORT`

Port that the GraphQL service listens on (`8080` by default).

#### `MONGODB_URL`

Address of the MongoDB instance to use (e.g. `mongodb://localhost:27017`).

#### `NINKA_JWT_SECRET`

base64-encoded secret for JWT HMAC signing (e.g. the output of
`head -c128 /dev/urandom | base64`).

#### `NINKA_PORT`

Port that the Ninka service listens on (`8084` by default).

#### `NINKA_URL`

Address of the Ninka service (listening at `localhost:8084` by default).

#### `NINSHOU_PORT`

Port that the Ninshou service listens on (`8081` by default).

#### `NINSHOU_URL`

Address of the Ninshou service (listening at `localhost:8081` by default).

#### `SHINNINJOU_PORT`

Port that the Shinninjou service listens on (`8083` by default).

#### `SHINNINJOU_URL`

Address of the Shinninjou service (listening at `localhost:8083` by default).

#### `YOUJI_PORT`

Port that the Youji service listens on (`8082` by default).

#### `YOUJI_URL`

Address of the Youji service (listening at `localhost:8082` by default).

### Running

In separate terminal sessions:

```bash
go run ./cmd/graphql
go run ./cmd/ninka
go run ./cmd/ninshou
go run ./cmd/shinninjou
go run ./cmd/youji
```

### Updating GraphQL or Protocol Buffer Schemas

To regenerate everything:

```bash
go generate ./...
```

Or just the GraphQL or a single protobuf schema:

```bash
go generate ./api/graphql
```

## Copyright

Copyright © 2019 Jesse B. Hannah. Licensed under the [GNU AGPL version 3 or
later][agpl].

[agpl]: LICENSE
[go]: https://golang.org/
[mkcert]: https://github.com/FiloSottile/mkcert
[mongodb]: https://www.mongodb.com/
