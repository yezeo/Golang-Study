package main

// Scanln()을 이용하여 숫자 2개를 입력받는 예제

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(a, b)
	}
}
