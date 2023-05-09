package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", ":8080")
	var reply = new(string)
	err := client.Call("HelloService.Hello", "huang", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
