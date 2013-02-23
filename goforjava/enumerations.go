package main

import (
	"fmt"
)

// START OMIT
const (
	SUN = iota
	MON
	TUE
	WED
	THU
	FRI
	SAT
)

// END OMIT

func main() {
	fmt.Println(SUN, MON, TUE, WED, THU, FRI, SAT)
}
