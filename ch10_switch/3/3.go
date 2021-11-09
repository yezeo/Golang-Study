package main

import "fmt"

func main() {

	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날입니다.")
		fmt.Println("둘쨰 날입니다.")
	case 2:
		fmt.Println("둘째 날입니다.")
		fmt.Println("새로운 팀원 면접이 있습니다")
	default:
		fmt.Println("프로젝트를 진행하세요.")
	}

}
