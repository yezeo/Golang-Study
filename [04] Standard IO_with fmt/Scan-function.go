package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b) // n과 err에 입력 두 값 받기
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n,a,b)
	}

}
