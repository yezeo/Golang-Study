package main

import "fmt"

func main() {
	var a = 123.12345
	var b = 3.14

	// %f와 %g 차이 비교
	fmt.Printf("%08.2f\n", a) //최소 너비 8, 소수점 이하 2자리, 공백은 0으로 채움
	fmt.Printf("%08.2g\n", a) //최소 너비 8, 총 숫자 2자리, 공백은 0으로 채움

	fmt.Printf("%8.5g\n", a) //최소 너비 8, 총 숫자 5자리
	fmt.Printf("%f\n", b) // 소수점 이하 6자리까지 출력 

}
