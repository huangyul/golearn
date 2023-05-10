package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn_go/grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	res, _ := c.SayHello(context.Background(), &proto.HelloRequest{Name: "huang"})
	fmt.Println(res.Message)
}
