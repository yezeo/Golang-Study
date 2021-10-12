package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)    // %5d
	fmt.Printf("%05d, %05d\n", a, b)  // %05d
	fmt.Printf("%-5d, %-05d\n", a, b) // %-5d (왼쪽 정렬)

	// 지정한 최소 너비보다 긴 값일 경우
	fmt.Printf("%5d\n", c)
	fmt.Printf("%05d\n", c)
}
