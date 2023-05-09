package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cupInfo(ctx context.Context) {
	defer wg.Done()
	fmt.Println("traceid: %s\r\n", ctx.Value("traceid"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println("over")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("cpu info")
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "traceid", "123123")
	wg.Add(1)
	go cupInfo(ctx2)
	time.Sleep(time.Second * 6)
	// 由父方法取消监听
	cancel()
	wg.Wait()
	fmt.Println("success")

}
