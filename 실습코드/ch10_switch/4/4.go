package main

import "fmt"

func main() {

	day := "Thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월화요일은 수업하러 가는 날이다")
	case "Wednesday", "Thursday":
		fmt.Println("수목금요일은 실습하는 날이다")
	}
}
