This repository is for reproducing a bug that race option doesn't work with binay-only package

## How to test
Clone the repository and run the below under the repository

```
GOPATH=$(go env GOPATH)
go build -o $GOPATH/pkg/darwin_amd64/github.com/motomux/hellobop.a -x ./src
go clean -cache
go test ./test
go test -race ./test
```
