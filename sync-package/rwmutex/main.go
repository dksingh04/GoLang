package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Since map is not thread safe hence while writing or reading data to or from
map we need to make it thread safe and we can make the map the thread safe by using RWMutex API form
sync package.
*/
type mapCounter struct {
	m map[int]int
	sync.RWMutex
}

var wg sync.WaitGroup

func main() {
	mc := &mapCounter{m: make(map[int]int)}
	wg.Add(1)
	go writeToMap(mc, 5)
	wg.Add(1)
	go readFromMap(mc, 5)
	wg.Add(1)
	go readFromMap(mc, 5)

	wg.Wait()

	fmt.Println("Main Go Routine exitted")
}

func writeToMap(mc *mapCounter, n int) {
	mc.Lock()
	defer wg.Done()
	for i := 0; i < n; i++ {
		mc.m[i] = i * (i + 1)
	}
	mc.Unlock()
}

func readFromMap(mc *mapCounter, n int) {
	mc.RLock()
	defer wg.Done()
	v := mc.m[rand.Intn(n)]
	mc.RUnlock()
	fmt.Println(v)
}
