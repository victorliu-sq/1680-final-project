install protoc and gRPC

```shell
apt install -y protobuf-compiler

go get google.golang.org/grpc
go 
```



error

```shell
# ERROR: protoc-gen-go-grpc: unable to determine Go import path for

# SOLUTION: add this line to rpcMsg.proto
option go_package = "pkg/protocl/rpcMSg";

# ERROR --go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1

# SOLUTION
export PATH="$PATH:$(go env GOPATH)/bin"
```

