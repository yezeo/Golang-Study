package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // 초기문; 조건문; 후처리
		fmt.Print(i, ", ") //i 값 출력
	}

	//  fmt.Println(i)              // Error - i는 이미 사라졌다.
}
