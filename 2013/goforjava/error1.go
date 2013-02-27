package main

import (
	"fmt"
)

// START OMIT
type FileError struct {
}

func (this *FileError) Error() string { // HL
	return "FileError"
}

func FileOpen(path string) error { // HL
	if "" == path {
		return &FileError{}
	}
	return nil
}

func main() {
	if err := FileOpen(""); nil != err { // HL
		fmt.Println(err)
		return
	}
	fmt.Println("bye")
}

// END OMIT
