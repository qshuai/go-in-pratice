package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("start exec goroutine")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("receive stop signal")
				wg.Done()
				return
			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	cancelFunc()
	wg.Wait()

	fmt.Println("program stoped")
}
