package main

import (
	"fmt"
)

type square struct {
	sidelength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {

	sq := square{}
	tr := triangle{}
	sq.sidelength = 5.0

	tr.base = 3.0
	tr.height = 4.0

	printArea(sq)
	printArea(tr)

}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
