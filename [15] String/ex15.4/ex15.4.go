package main

import "fmt"

func main() {
	str1 := "가나다라마"
	str2 := "abcde"

	fmt.Printf("'가나다라마'의 크기= %d\n", len(str1)) //한글 문자열 크기
	fmt.Printf("'abcde'의 크기= %d\n", len(str2)) //영문 문자열 크기

}
