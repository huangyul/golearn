package main

import "fmt"

func mPrint(data ...interface{}) {
	for _, value := range data {
		fmt.Println(value)
	}
}

func main() {
	var data = []interface{}{
		"huang", 18, 1.80,
	}
	mPrint(data...)

}
