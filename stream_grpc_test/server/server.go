package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learn_go/stream_grpc_test/proto"
	"net"
	"strconv"
	"sync"
	"time"
)

const PORT = ":8080"

type server struct{}

func (s *server) GetStream(request *proto.StreamReqData, response proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		err := response.Send(&proto.StreamResData{Data: strconv.Itoa(i)})
		if err != nil {
			return err
		}
		i++
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PostStream(streamServer proto.Greeter_PostStreamServer) error {
	for {
		a, err := streamServer.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
	return nil
}

func (s *server) AllStream(streamServer proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		wg.Done()
		for {
			data, _ := streamServer.Recv()
			fmt.Println("收到客户端的消息" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = streamServer.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", PORT)
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &server{})
	_ = g.Serve(lis)

}
