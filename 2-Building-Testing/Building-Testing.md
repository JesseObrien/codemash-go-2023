# Building and Testing

## Wat

- wat is `go mod` even?
  - Dependency manager
  - Keeps things `tidy`
  - `go get` when to use it
    - Upgrading existing dependencies
- wat is `GO111MODULE=on`??
  - Prior to go1.11
  - $GOPATH/pkg/mod cache
- wat is `go clean`
  - Clean up any build rubbish
  - `go clean -testcache`
  - `go clean -modcache`

## Fast Builds Save Cycles

What do we want?! **Fast Builds!!**

When do we want them?! **Seconds Ago!!**

We have options:

- Reflex: `go install github.com/cespare/reflex@latest`
- Modd: `env GO111MODULE=on go install github.com/cortesi/modd/cmd/modd@latest`
