package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learn_go/grpc_metadata_test/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := proto.NewHelloServerClient(conn)

	// 定义metadata
	md := metadata.New(map[string]string{
		"name": "huang",
		"age":  "23",
	})
	// 生成一个新的ctx，带上md
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, _ := c.SayHello(ctx, &proto.HelloRequest{Name: "huang"})
	fmt.Println(res)
}
