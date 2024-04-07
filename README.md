# IM Chat Demo


安装protobuf库文件
```shell
go get github.com/golang/protobuf/proto
```

在根目录测试：
```shell
protoc --go_out=pkg/protocol/ pkg/protocol/*.proto
```
