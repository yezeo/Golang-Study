package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a) // 2바이트로 전환할 때 정수 int16에서 1바이트 정수 int8로 변환할 때 상위 1바이트가 없어짐.

	fmt.Println(a)
	fmt.Println(c)
}
