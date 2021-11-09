package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(angle float64) Direction {
	switch true {
	case angle >= 315:
		return North
	case angle < 45 && angle >= 0:
		return North
	case angle >= 45 && angle <= 135:
		return East
	case angle > 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return North
	default:
		return None
	}
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
