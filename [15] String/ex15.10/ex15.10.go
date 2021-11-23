package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2 // str1 + " " + str2 와 동일
	fmt.Println(str1)

}
