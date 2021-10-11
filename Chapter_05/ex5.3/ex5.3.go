package main

import "fmt"

func main() {
	var a = 324.123455
	var c = 3.14

	fmt.Printf("%08.2f\n", a) // 최소 너비 8, 소수점 이하 2자리, 0을 채움
	fmt.Printf("%08.2g\n", a) // 최소 너비 8, 총숫자 2자리, 0을 채움
	fmt.Printf("%8.5g\n", a)  // 최소 너비 8, 총숫자 5자리
	fmt.Printf("%f\n", c)     // 소수점 이하 6자리까지 출력
}
