package main

import "fmt"

func main() {
	str := "Hello World"
	runes := []rune(str) // 문자열을 []rune 타입으로 변환

	fmt.Printf("len(str) = %d\n", len(str))     // 메모리 크기 반환
	fmt.Printf("len(runes) = %d\n", len(runes)) //글자수 반환; len(runes)
}
