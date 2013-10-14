package main

import (
	"fmt"
	"time"
)

func SearchImage(done chan string) {
	var data string
	for i := 0; i < 10; i++ {
		fmt.Print(1, " ")
		data += "1 "
		time.Sleep(500 * time.Millisecond)
	}
	done <- data
}

func SearchBlog(done chan string) {
	var data string
	for i := 0; i < 10; i++ {
		fmt.Print(0, " ")
		data += "0 "
		time.Sleep(300 * time.Millisecond)
	}
	done <- data
}

// START OMIT
func main() {
	fmt.Println("start")
	defer fmt.Println("stop")

	c1 := make(chan string) // HL
	c2 := make(chan string) // HL

	go SearchBlog(c1)
	go SearchImage(c2)

	blogData, imageData := <-c1, <-c2 // HL
	fmt.Println()
	fmt.Println(blogData)
	fmt.Println(imageData)
}

// END OMIT
