package main

import (
	"context"
	"google.golang.org/grpc"
	"learn_go/grpc_test/proto"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	_ = g.Serve(lis)
}
