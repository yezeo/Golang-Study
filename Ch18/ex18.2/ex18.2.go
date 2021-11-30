package main

import "fmt"

func main() {

	var slice = []int{1, 2, 3}

	slice2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(slice2)
}
