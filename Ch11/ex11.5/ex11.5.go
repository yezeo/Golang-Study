package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { //5번 반복
		for j := 0; j < i+1; j++ { //현재 i값+1만큼 반복
			fmt.Print("*") // * 출력
		}
		fmt.Println()
	}
}
