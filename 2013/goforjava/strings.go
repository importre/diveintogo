package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var str string = "Go 프로그래밍"
	s := []rune(str) // HL
	for i, c := range s {
		fmt.Printf("%2d: %c\n", i, c)
	}
	// END OMIT
}
