package main

import "fmt"

func main() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Printf("%6d\n", a)
	fmt.Printf("%06d\n", b)
	fmt.Printf("%6.2f\n", f)
}
