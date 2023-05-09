package main

import "net/rpc"

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloServer", HelloService{})
}
