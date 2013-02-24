package main

import (
	"fmt"
)

func main() {
	// START OMIT
	a1 := [...]int{1, 2, 3}

	a2 := a1 // HL
	fmt.Println("a1 == a2:", a1 == a2)

	a1[0] = 100
	fmt.Println("a1 == a2:", a1 == a2)

	fmt.Println(a1)
	fmt.Println(a2)
	// END OMIT
}
