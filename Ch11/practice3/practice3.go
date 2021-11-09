package main

import "fmt"

func main() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i, "*", i, "=", i*i)
	}
}
