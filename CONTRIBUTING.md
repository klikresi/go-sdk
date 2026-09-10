# Contributing

Thanks for helping improve the Klik Resi Go SDK!

## Setup

```sh
git clone git@github.com:klikresi/go-sdk.git
cd go
go mod download
```

## Checks

Run the following before opening a pull request:

```sh
gofmt -l .          # must print nothing
go vet ./...
go test -race ./...
```

## Tests

Tests run entirely offline against recorded API responses stored in
`testdata/`. No API key is required.

## Conventions

- Follow standard Go idioms and keep the public API surface small.
- Every exported symbol must have a doc comment.
- No external dependencies.
