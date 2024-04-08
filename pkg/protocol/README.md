
Install protobuf
```shell
go get github.com/golang/protobuf/proto
```

In root directory execute
```shell
protoc --go_out=pkg/protocol/ pkg/protocol/*.proto
```
