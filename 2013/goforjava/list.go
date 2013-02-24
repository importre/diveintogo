package main

// START OMIT
import (
	"container/list" // HL
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack("GDG Korea")
	l.PushBack("Golang")
	l.PushBack(1)

	for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
	}
}

// END OMIT
