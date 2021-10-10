package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // c는 3
	d := float64(a * c) // d는 9.0

	var e int64 = 7
	f := int64(d) * e // f는 63

	var g int = int(b * 3) // g는 10
	var h int = int(b) * 3 // h는 9
	fmt.Println(g, h, f)
}
