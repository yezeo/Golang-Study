package main

import "fmt"

func main() {
	var a = 123
	var b int = 4567
	f := 3.1459269

	fmt.Printf("%6d", a)
	fmt.Printf("%06d", b)
	fmt.Printf("%6.2f", f)
}
