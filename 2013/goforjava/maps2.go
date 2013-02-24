package main

import (
	"fmt"
)

type any interface{}

func main() {
	m := make(map[any]bool)

	// START OMIT
	m[1.0] = true
	m["go"] = true
	m[1] = false

	for k, v := range m { // HL
		fmt.Println(k, v)
	}
	fmt.Println()

	delete(m, 1) // HL

	for k, v := range m { // HL
		fmt.Println(k, v)
	}
	// END OMIT
}
