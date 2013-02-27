package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var f float64
	i := 10

	// f = float64(i)
	f = i // HL
	fmt.Printf("%T\n", f)
	// END OMIT
}
