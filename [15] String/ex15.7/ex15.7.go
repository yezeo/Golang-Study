package main

import "fmt"

func main() {
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T\t값:%d\t문자값:%c\n", str[i], str[i], str[i])
	}
}
