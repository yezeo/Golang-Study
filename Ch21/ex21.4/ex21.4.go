package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {

		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {

		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4)
	fmt.Println(result)
}
