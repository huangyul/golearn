package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learn_go/grpc_metadata_test/proto"
	"net"
)

type Server struct {
}

func (receiver Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	for key, val := range md {
		fmt.Println(key, val)
	}
	return &proto.HelloResponse{Message: "hello, " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterHelloServerServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	g.Serve(lis)
}
