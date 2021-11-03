package main

import "fmt"

func F(n int) int {
	if n < 2 {
		return n
	}
	return F(n-2) + F(n-1)
}

func main() {
	fmt.Println(F(9))
}
