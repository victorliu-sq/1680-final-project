all:
	go build cmd/control/snowcast_control.go
	go build cmd/listener/snowcast_listener.go
	go build cmd/server/snowcast_server.go

clean:
	rm snowcast_control snowcast_listener snowcast_server