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

	c := make(chan string) // HL

	go SearchBlog(c)
	go SearchImage(c)

	blogData, imageData := <-c, <-c // HL
	fmt.Println()
	fmt.Println(blogData)
	fmt.Println(imageData)
}

// END OMIT
