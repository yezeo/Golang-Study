package main

import "fmt"

var g int = 10

func main() {
	var m int = 20
	{
		var s int = 50
		fmt.Println(m, s, g)
	}
	m = s + 20 // error
	// .\ex4.6.go:13:6: undefined: s
}
