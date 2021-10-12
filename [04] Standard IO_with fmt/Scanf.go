package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b) // n과 err에 입력 두 값 받기
	if err != nil {            // 실패한 경우
		fmt.Println(n, err)
	} else { // 성공한 경우
		fmt.Println(n, a, b)
	}

}
