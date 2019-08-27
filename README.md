# Golang Algorithms

## build

on Windows:

```cmd
set GOPROXY=https://goproxy.io
go test ./...
go test -cover ./...
go test -bench=./...
go build -o golang-algorithms.exe .
```

or Linux:

```shell
export GOPROXY=https://goproxy.io
go test ./...
go test -cover ./...
go test -bench=./...
go build -o golang-algorithms .
```
