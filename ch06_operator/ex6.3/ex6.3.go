package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2:%08b x<<2:%d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2:%08b y<<2:%d\n", y, y<<2, y<<2)
}
