package main

import (
	"fmt"
)

func main() {
	// START OMIT
	b := [5]byte{1, 2, 3, 4, 5} // array // HL

	fmt.Printf("b  : %v\ncap: %v\nlen: %v\n\n", b, len(b), cap(b))

	s := b[2:4] // slice // HL

	fmt.Printf("s  : %v\ncap: %v\nlen: %v\n", s, len(s), cap(s))
	// END OMIT
}
