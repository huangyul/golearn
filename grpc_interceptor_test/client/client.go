package main

import (
	"context"
	"google.golang.org/grpc"
	"learn_go/grpc_interceptor_test/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	c := proto.NewGreeterClient(conn)
	c.SayHello(context.Background(), &proto.HelloRequest{Name: "huang"})
}
