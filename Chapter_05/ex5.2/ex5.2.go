package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d, %-5d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-5d\n", c, c)
}
