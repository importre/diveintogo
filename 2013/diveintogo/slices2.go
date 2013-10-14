package main

import (
	"fmt"
)

func pushBack(slice []int, item ...int) []int {
	return append(slice, item...)
}

func main() {
	// START OMIT
	a1 := [20]int{}
	s1 := a1[:10]
	s2 := a1[10:15]
	s3 := a1[:]

	a1[11] = 100
	s2 = pushBack(s2, 1, 2, 3, 4, 5)

	fmt.Printf("len: %2d, cap: %2d\n%v\n\n", len(s1), cap(s1), s1)
	fmt.Printf("len: %2d, cap: %2d\n%v\n\n", len(s2), cap(s2), s2)
	fmt.Printf("len: %2d, cap: %2d\n%v\n\n", len(s3), cap(s3), s3)
	// END OMIT
}
