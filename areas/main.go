package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{
		base:   15.41,
		height: 13.1,
	}

	sq := square{
		sideLength: 32.1,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}
