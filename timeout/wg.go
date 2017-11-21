package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// try again less than three times
	wg.Add(3)

	//mock timeout by create an empty channel
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("process success")
			case <-time.After(1 * time.Second):
				fmt.Println("try again")
				wg.Done()
			}
		}
	}()

	wg.Wait()
}
