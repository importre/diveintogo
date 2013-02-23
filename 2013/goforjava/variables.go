package main

import (
	"fmt"
)

func main() {
    // START OMIT
	var i = 10 // var keyword // HL

	var j int // var keyword with int type // HL
	j = 11

	k := 12 // init operator // HL

	fmt.Println(i, j, k)

	a, b, c := "고", 123.1, &i // multiple variables // HL

	fmt.Printf("%v %v %v\n", a, b, c)
	fmt.Printf("%T %T %T\n", a, b, c)
    // END OMIT
}
