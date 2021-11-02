package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 5)
	fmt.Println(c)
}
