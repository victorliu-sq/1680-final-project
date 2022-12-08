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

runC1:
	./snowcast_control localhost 8888 5000 control1

runL:
	./snowcast_listener 5000 | pv > /dev/null

runS:
	./snowcast_server 8888 \
	./mp3/Beethoven-SymphonyNo5.mp3 ./mp3/DukeEllington-Caravan.mp3 ./mp3/FX-Impact193.mp3