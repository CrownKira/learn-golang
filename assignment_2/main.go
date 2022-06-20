package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	height float64
}

type square struct {
	sideLength float64
	sideHeight float64
}

func main() {
	t := triangle {
		base: 5,
		height: 6,
	}
	s := square {
		sideLength: 7,
		sideHeight: 8,
	}
	printArea(t)
	printArea(s)		
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideHeight
}