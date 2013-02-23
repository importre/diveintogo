package main

import (
	"fmt"
)

func main() {
	// START OMIT
    for i, c := range "Go 프로그래밍" { // HL
        fmt.Printf("%2d: 0x%04X %c\n", i, c, c)
    }
	// END OMIT
}
