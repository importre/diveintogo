package main

import (
	"fmt"
)

// START OMIT
func IsHangul(c rune) bool {
	var ret bool
	switch {
	case '가' <= c && c <= '힣':
		ret = true
	case 'a' <= c && c <= 'z':
		ret = false
	}
	return ret
}

func main() {
	for _, c := range "a고" {
		fmt.Printf("%c\t%t\n", c, IsHangul(c))
	}
}

// END OMIT
