package main

import (
	"fmt"
)

// START OMIT
type Integer int // HL

func (self Integer) FloatValue() float64 {
	return float64(self)
}

func main() {
	i := Integer(10)
	fmt.Printf("type: %T\n", i)
	fmt.Printf("type: %T\n", i.FloatValue())
}

// END OMIT
