package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)
import "learn_go/grpc_interceptor_test/proto"

type Server struct{}

func (receiver Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "hello, " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 拦截器的逻辑
		fmt.Println("receipt new request")
		// 执行以前的逻辑，相当于next
		res, err := handler(ctx, req)
		fmt.Println("请求完成")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)
	lis, _ := net.Listen("tcp", ":8080")
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	err := g.Serve(lis)
	if err != nil {
		return
	}

}
