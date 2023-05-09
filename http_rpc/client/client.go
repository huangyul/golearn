package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply = new(string)
	client.Call("HelloService.Hello", "huang", reply)
	fmt.Println(*reply)
}
