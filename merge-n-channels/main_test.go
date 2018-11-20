package main

import "testing"

var funcs = []struct {
	name string
	f    func(...<-chan int) <-chan int
}{
	{"goroutines", merge},
	{"reflection", mergeReflect},
	{"recursive", mergeRecursive},
}

//TODO Test Merge functionality
func TestMerge(t *testing.T) {

}

//TODO Benchmark Test of different Merge
func BenchMarkMerge(b *testing.B) {

}
