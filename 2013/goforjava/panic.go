package main

import (
	"fmt"
)

// START OMIT
func Div(a, b float64) float64 {
	if 0.0 == b {
		panic("Don't divide by zero") // HL
	}
	return a / b
}

func Cal() {
	defer func() { // HL
		if r := recover(); r != nil { // HL
			fmt.Println("Recovered in Cal():", r)
		}
	}()

	a := Div(12.0, 0.0)
	fmt.Println(a)
}

// END OMIT

func main() {
	Cal()
}
