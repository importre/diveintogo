package main

import (
	"fmt"
	"os"
)

// START OMIT
func Open(path string) bool {
	file, err := os.Open(path)

	if nil != err {
		fmt.Println(err)
		return false
	}

	defer file.Close() // HL

	// do something

	if nil != err {
		fmt.Println(err)
		return false
	}

	return true
}

// END OMIT

func main() {
	a := Open("1.txt")
	fmt.Println(a)
}
