package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a // a의 메모리 공간
	var p2 *int = &a // a의 메모리 공간
	var p3 *int = &b // b의 메모리 공간

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
}
