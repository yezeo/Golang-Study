package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i, "*", i, "=", i*i)
		}
	}
}
