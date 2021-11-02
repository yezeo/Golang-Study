package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
}

func main() {
	// math := 80
	// eng := 74
	// history := 95
	// fmt.Println("김일동 님 평균 점수는", (math+eng+history)/3, "입니다.")
	PrintAvgScore("김일동", 80, 74, 95)

	// math = 80
	// eng = 73
	// history = 78

	// fmt.Println("박상동 님 평균 점수는", (math))
	PrintAvgScore("박상동", 80, 73, 78)
	PrintAvgScore("송이동", 78, 73, 78)
}
