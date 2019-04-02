# issho-ni

[![Go Report Card](https://goreportcard.com/badge/github.com/issho-ni/issho)](https://goreportcard.com/report/github.com/issho-ni/issho)

## Requirements

- [Go][] 1.12
- [MongoDB][]
- [mkcert][]

## Development

```bash
export GO111MODULE=on
mkcert localhost 127.0.0.1 ::1
```

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
