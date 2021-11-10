package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ { // 3번 반복
		for j := 0; j < 5; j++ { //5번 반복
			fmt.Print("*") // * 출력
		}
		fmt.Println()
	}
}
