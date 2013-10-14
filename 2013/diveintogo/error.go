package main

import (
	"fmt"
)

// START OMIT
type FileError struct {
	path string
}

func (this *FileError) Error() string { // HL
	if "" == this.path {
		return "path error"
	}
	return ""
}

func FileOpen(path string) error { // HL
	return &FileError{path}
}

func main() {
	if err := FileOpen(""); nil != err {
		fmt.Println(err)
		return
	}
}

// END OMIT
