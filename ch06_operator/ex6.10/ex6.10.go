package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	a, b = b, a

	fmt.Println(a, b)
}
