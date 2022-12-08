all:
	go build cmd/control/snowcast_control.go
	go build cmd/listener/snowcast_listener.go
	go build cmd/server/snowcast_server.go

clean:
	rm snowcast_control snowcast_listener snowcast_server

PROTO_DIR = pkg/protocol/rpcMsg/

generate:
	protoc -I $(PROTO_DIR) \
	--go_out=. \
    --go-grpc_out=. \
    $(PROTO_DIR)rpcMsg.proto