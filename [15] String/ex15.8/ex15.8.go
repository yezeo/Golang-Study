package main

import "fmt"

func main() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T\t값:%d\t문자값:%c\n", arr[i], arr[i], arr[i])
	}
}
