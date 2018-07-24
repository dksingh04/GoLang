package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// make sure to call cancel after end of the task, to avoid leakage
	defer cancel()

	for n := range generate(ctx) {
		if n == 5 {
			//cancel()
			break
		}
	}
	time.Sleep(10 * time.Second)
}

func generate(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		n = 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Task is Done...")
				return
			case ch <- n:
				n++
				fmt.Printf("Task... %d is done\n", n)
			}
		}
	}()

	return ch
}
