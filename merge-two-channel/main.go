package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a := asChan(1, 2, 3, 4, 5)
	b := asChan(6, 7, 8, 9, 10)

	c := merge(a, b)
	//go func() {
	for v := range c {
		fmt.Println(v)
	}

	//}()

}

func merge(a, b <-chan int) <-chan int {

	c := make(chan int)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					fmt.Println("a is done")
					continue
				}
				c <- v
			case v, ok := <-b:

				if !ok {
					b = nil
					fmt.Println("b is done")
					continue
				}

				c <- v

			}
		}
	}()

	return c
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}
