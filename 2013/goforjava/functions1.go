package main

import (
	"fmt"
)

// START OMIT
func add(i int, j int) (k int) {
	fmt.Printf("k = %v\n", k) // initialized to the zero value
	k = i + j
	return // just return // HL
}

// END OMIT

func main() {
	i := add(1, 2)
	fmt.Println(i)
}
