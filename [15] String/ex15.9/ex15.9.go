package main

import "fmt"

func main() {
	str := "Hello 월드!"
	for _, v := range str {
		fmt.Printf("타입: %T\t값:%d\t문자값:%c\n", v, v, v)
	}
}
