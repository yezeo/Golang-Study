package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바깥활동에 좋은 날씨가 아닙니다")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하시죠")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다")
	default:
		fmt.Println("따듯합니다")
	}
}
