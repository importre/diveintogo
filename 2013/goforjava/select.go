package main

import (
	"fmt"
	"time"
)

func SearchImage(done chan string) {
	var data string
	for i := 0; i < 10; i++ {
		data += "1 "
		time.Sleep(500 * time.Millisecond)
	}
	done <- data
}

func SearchBlog(done chan string) {
	var data string
	for i := 0; i < 10; i++ {
		data += "0 "
		time.Sleep(300 * time.Millisecond)
	}
	done <- data
}

func main() {
	fmt.Println("start")
	defer fmt.Println("stop")

	// START OMIT
	c1, c2 := make(chan string), make(chan string) // HL
	done := make(chan bool, 2)

	go func() {
		for {
			select { // HL
			case a := <-c1:
				fmt.Println(a)
				done <- true
			case b := <-c2:
				fmt.Println(b)
				done <- true
			}
		}
	}()

	go SearchBlog(c1)
	go SearchImage(c2)

	for i := 0; i < 2; i++ {
		fmt.Println(<-done)
	}
	// END OMIT
}
