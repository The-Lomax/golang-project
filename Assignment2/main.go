package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	tr := triangle{
		base:   6,
		height: 5,
	}
	sq := square{
		sideLength: 5,
	}

	printArea(tr)
	printArea(sq)
}
