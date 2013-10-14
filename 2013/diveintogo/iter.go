package main

import (
	"fmt"
	"time"
)

// START OMIT
func dump(arr []int) {
	for index, value := range arr {
		fmt.Println(index, value)
	}
	fmt.Println()
}

func main() {
	elapsedTime := func(arr []int, dump func(arr []int)) {
		beg := time.Now()
		dump(arr)
		end := time.Now()
		fmt.Println(end.Sub(beg))
	}

	a := [...]int{10, 9, 8, 7, 6}
	elapsedTime(a[:], dump)
}
// END OMIT
