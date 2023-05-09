package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 实例化一个server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("server初始化失败")
	}
	// 注册处理逻辑
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
