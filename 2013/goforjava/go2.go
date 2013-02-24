package main

import (
	"fmt"
	"time"
)

func SearchImage() {
	var data string
	for i := 0; i < 10; i++ {
		fmt.Print(1, " ")
		data += "1 "
		time.Sleep(500 * time.Millisecond)
	}
}

func SearchBlog() {
	var data string
	for i := 0; i < 10; i++ {
		fmt.Print(0, " ")
		data += "0 "
		time.Sleep(300 * time.Millisecond)
	}
}

// START OMIT
func main() {
	fmt.Println("start")
	defer fmt.Println("stop")

	go SearchBlog() // HL
	go SearchImage() // HL
	time.Sleep(10 * time.Second)
	fmt.Println()
}

// END OMIT
