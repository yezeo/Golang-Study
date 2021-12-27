package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(30))
	var b int = 20
	fmt.Println(myInt(b).add(50))
}
