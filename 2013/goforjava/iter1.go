package main

import (
	"fmt"
	"time"
)

// START OMIT
func dump1(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func dump2(arr []int) {
	for _, v := range arr {
		fmt.Print(v)
	}
	fmt.Println()
}

// END OMIT

func main() {
	elapsedTime := func(arr []int, dump func(arr []int)) {
		beg := time.Now()
		dump(arr)
		end := time.Now()
		fmt.Println(end.Sub(beg))
	}

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	elapsedTime(a[:], dump1)
	elapsedTime(a[:], dump2)
}
