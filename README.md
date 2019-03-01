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

### Updating GraphQL Schema

```bash
go gen ./...
```

## Copyright

Copyright © 2019 Jesse B. Hannah. Licensed under the [GNU AGPL version 3 or
later][agpl].

[agpl]: LICENSE
[go]: https://golang.org/
[mkcert]: https://github.com/FiloSottile/mkcert
