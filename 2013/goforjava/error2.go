package main

import (
	"errors" // HL
	"log"
)

func Check(path string) error {
	if "" == path {
		return errors.New("param error") // HL
	}
	return nil
}

func main() {
	if err := Check(""); nil != err {
		log.Fatal(err)
	}

	log.Println("bye")
}
