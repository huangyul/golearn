package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8080/%s?a=%d&b=%d"), "add", a, b)
	body, _ := res.Body()
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	fmt.Println(Add(2, 4))
}
