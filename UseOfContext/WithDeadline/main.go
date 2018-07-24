package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	duration := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), duration)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Overselept!!")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
