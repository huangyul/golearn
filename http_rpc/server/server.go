package main

import (
	"learn_go/http_rpc/handler"
	"learn_go/http_rpc/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	//_ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
	server_proxy.RegisterHelloService(&handler.HelloService{})
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

}
