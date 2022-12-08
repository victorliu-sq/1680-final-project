install protoc and gRPC

```shell
apt install -y protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```



error

```shell
#protoc-gen-go-grpc: unable to determine Go import path for

# add this line to rpcMsg.proto
option go_package = "./";

#
```

