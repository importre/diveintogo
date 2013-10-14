package main

import (
	"fmt"
)

// START OMIT
type Point struct {
	x, y float64
}

type Segment struct {
	p1, p2 Point
}

func (self *Point) ToString() string {
	return fmt.Sprintf("x: %v, y: %v", self.x, self.y)
}

func (self *Segment) ToString() string {
	return fmt.Sprintf("p1 %v\np2 %v\n", self.p1.ToString(), self.p2.ToString())
}

type Shaper interface { // HL
	ToString() string
}

// END OMIT

func main() {
	var shaper Shaper

	shaper = &Point{1, 2}
	fmt.Println(shaper.ToString())

	shaper = &Segment{Point{3, 4}, Point{5, 6}}
	fmt.Println(shaper.ToString())
}
