package main

import (
	"fmt"
	"sync"
)

type safeOperation struct {
	i int
	sync.Mutex
}

var wg sync.WaitGroup

func main() {
	so := new(safeOperation)
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go so.Increment(ch)
		wg.Add(1)
		go so.Decrement(ch)
	}

	go func() {
		for c := range ch {
			fmt.Println("Receiving Channel Value: ", c)
			wg.Done()
		}
	}()
	wg.Wait()
	//<-done
	fmt.Println("Value: ", so.GetValue())
	fmt.Println("Main method finished")
}

func (so *safeOperation) Increment(ch chan int) {
	so.Lock()
	so.i++
	ch <- so.i
	so.Unlock()
}

func (so *safeOperation) Decrement(ch chan int) {
	so.Lock()
	so.i--
	ch <- so.i
	so.Unlock()
}

func (so *safeOperation) GetValue() int {
	so.Lock()
	v := so.i
	so.Unlock()
	return v
}
