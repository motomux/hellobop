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

## Result with Go 1.10.3
Test is failed with `-race` option 
```
go test ./test
ok      github.com/motomux/hellobop/test        0.006s
```

```
go test -race ./test
go build github.com/motomux/hellobop: missing or invalid binary-only package
FAIL    github.com/motomux/hellobop/test [build failed]
```
