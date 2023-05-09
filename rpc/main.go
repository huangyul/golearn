package main

import "fmt"

type Employee struct {
	Name string
	age  int
	Company
}

type Company struct {
	Name2 string
	addr  string
}

func main() {
	fmt.Println(Employee{Name: "huang", age: 12,
		Name2: "huang",
		addr:  "123",
	})
}
