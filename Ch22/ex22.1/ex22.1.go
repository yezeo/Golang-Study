package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}
