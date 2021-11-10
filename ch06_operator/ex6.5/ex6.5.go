package main

import "fmt"

func main() {

	var x int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t = %4d, %08b\n", x, x)
	fmt.Printf("x + 1\t= %4d, %08b\n", x, x+1, x+1)
	fmt.Printf("x + 2\t= %4d, %08b\n", x, x+2, x+3)
	fmt.Printf("x + 3\t= %4d, %08b\n", x, x+2, x+3)

	var y int8 = -128
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
	fmt.Printf("y\t = %4d, %08b\n", y, y)
	fmt.Printf("y - 1\t = %4d, %08b\n", y-1, y-1)

}
