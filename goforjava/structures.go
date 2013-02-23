package main

import (
	"fmt"
)

// START OMIT
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

// END OMIT

func (self *Point) X() float64 {
	return self.x
}

func main() {
	p := NewPoint(1.1, 2)
	fmt.Println(p)
	fmt.Println(p.X())
}
