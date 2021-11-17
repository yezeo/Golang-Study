package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.0, 30.0}

	for i, v := range t {
		fmt.Println(i, v)
	}

	for i, v := range t {
		fmt.Printf("t[%d] = %f\n", i, v) // v[0] = 24.0 이렇게 출력
	}
}
