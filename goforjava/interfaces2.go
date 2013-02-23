package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}

type Segment struct {
	p1, p2 Point
}

func (self *Point) ToString() string {
	return fmt.Sprintf("x: %v\ny: %v", self.x, self.y)
}

func (self *Segment) ToString() string {
	return fmt.Sprintf("p1 %v\np2 %v\n", self.p1.ToString(), self.p2.ToString())
}

type Shaper interface { // HL
	ToString() string
}

// START OMIT
func main() {
	var shaper Shaper

	shaper = &Point{}
	fmt.Println(shaper)

	shaper = &Segment{}
	fmt.Println(shaper)
}

// END OMIT
