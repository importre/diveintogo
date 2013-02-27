package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var a1 [2]int
	var a2 [2][2]bool = [2][2]bool{
		[2]bool{true, false},
		[2]bool{false, true},
	}

	a3 := [...]int{1, 2, 3} // is equivalent to "a3 := [3]int{1, 2, 3}
	a4 := [3]*int{}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	// END OMIT
}
