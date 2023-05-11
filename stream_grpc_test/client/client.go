package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn_go/stream_grpc_test/proto"
	"strconv"
	"sync"
	"time"
)

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithInsecure())
	defer conn.Close()
	client := proto.NewGreeterClient(conn)

	// 服务端流模式
	stream, err := client.GetStream(context.Background(), &proto.StreamReqData{Data: "huang"})
	if err != nil {
		return
	}
	for {
		a, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(a)
	}

	// 客户端流模式
	PStream, _ := client.PostStream(context.Background())
	i := 0
	for {
		i++
		PStream.Send(&proto.StreamReqData{Data: strconv.Itoa(i)})
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
	}

	// 双向流模式
	AllStream, _ := client.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := AllStream.Recv()
			fmt.Println("收到服务器的数据:" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := AllStream.Send(&proto.StreamReqData{Data: "来自客户端的信息"})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
