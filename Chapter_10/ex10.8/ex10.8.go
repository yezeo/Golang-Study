package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

// 각 ColorType 열거값에 따른 문자열을 반환하는 함수
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
