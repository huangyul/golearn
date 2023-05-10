package main

import (
	"fmt"
	"learn_go/http_rpc/client_proxy"
)

func main() {
	//client, _ := rpc.Dial("tcp", ":8080")
	HelloClient := client_proxy.NewHelloServiceClient("tcp", ":8080")
	var reply = ""
	//client.Call(handler.HelloServiceName+".Hello", "huang", &reply)
	HelloClient.Hello("huang", &reply)
	fmt.Println(reply)
}
