package main

import (
	"fmt"
)

// START OMIT
func printf(format string, args ...interface{}) (int, error) { // HL
	numOfBytes, err := fmt.Printf(format, args...)
	return numOfBytes, err
}

func main() {
	cnt := 0

	closure := func(str string) { // HL
		cnt++
		size, err := printf("%d: %s\n", cnt, str)
		fmt.Println(size, err)
	}

	closure("Hello")
	closure("Golang")
}
// END OMIT
