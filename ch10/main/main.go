package main

import (
	"fmt"
	"learn_go/ch10/user"
)

func main() {
	u := user.User{
		Name: "huang",
	}
	fmt.Println(user.GetUser(u))
}
