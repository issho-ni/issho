# issho-ni

## Requirements

- [Go][] 1.12
- [mkcert][]

## Development

```bash
export GO111MODULE=on
mkcert localhost 127.0.0.1 ::1
go run ./cmd/graphql
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
