package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false //제수가 0일 때 반환
	}
	return a / b, true //제수가 0이 아닐 때 반환
}

func main() {
	c, success := Divide(9, 3) //제수가 0이 아닌 경우
	fmt.Println(c, success)
	d, success := Divide(9, 0) //제수가 0인 경우
	fmt.Println(d, success)
}
