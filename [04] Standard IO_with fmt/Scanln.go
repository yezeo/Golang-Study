package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b) // 입력 두 값 받기- n은 성공 개수, err는 에러 반환
	if err != nil {            // 실패한 경우
		fmt.Println(n, err)
	} else { // 성공한 경우
		fmt.Println(n, a, b)
	}

}
