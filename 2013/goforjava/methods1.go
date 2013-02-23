package main

import (
	"fmt"
)

type Point struct { // HL
	x, y float64
}

func NewZeroPoint() *Point {
	return &Point{} // is equivalent to "return new(Point)"
}

func NewPoint(x, y float64) *Point {
	return &Point{
		y: y,
		x: x,
	} // is equivalent to "return &Point{x, y}"
}

// START OMIT
func (self *Point) X() float64 {
	return self.x
}

func main() {
	p := NewPoint(1.1, 2)
	fmt.Println("p  :", p)
	fmt.Println("p.x:", p.X())
}

// END OMIT
