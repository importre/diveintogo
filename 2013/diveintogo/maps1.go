package main

import (
	"fmt"
)

// START OMIT
type any interface{}

func main() {
	m := make(map[any]bool) // HL
	m[1.0] = true
	m["go"] = true

	res := m[1.0]
	fmt.Println(res)
	res = m["go"]
	fmt.Println(res)
	res = m["programming"] // zero value: false
	fmt.Println(res)
}

// END OMIT
