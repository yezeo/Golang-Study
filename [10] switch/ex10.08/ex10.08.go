package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota //0부터 시작해 1씩 증가
	Blue
	Green
	Yellow
)

// 각 ColorType 열거값에 따라 문자열 반환하는 함수
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
	return Green
}
func main() {
	fmt.Println("My Favorite color is", colorToString(getMyFavoriteColor()))
}
