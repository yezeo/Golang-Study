package main

import "fmt"

func main() {

	a := [2][5]int{ //이중배열 a 선언
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, //여기 참고- 닫는 중괄호가 같은 줄이 아니면 쉼표 찍기
	}
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ") // v값 출력
		}
	}
}
