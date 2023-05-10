package client_proxy

import (
	"learn_go/http_rpc/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol string, addr string) HelloServiceStub {
	conn, err := rpc.Dial(protol, addr)
	if err != nil {
		panic(err)
	}
	return HelloServiceStub{Client: conn}
}

func (s *HelloServiceStub) Hello(request string, reply *string) error {
	err := s.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
