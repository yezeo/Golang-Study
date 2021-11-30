package main

import "fmt"

func main() {
	var ch1 rune = '음'

	fmt.Printf("%T\n", ch1) // ch1의 타입 출력
	fmt.Println(ch1)        // ch1의 값 출력
	fmt.Printf("%c\n", ch1) // ch1에 저장된 값을 문자로 출력
}
