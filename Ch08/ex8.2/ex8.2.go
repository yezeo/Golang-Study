package main

import "fmt"

func main() {
	const PI1 float64 = 3.141592653589793238
	var PI2 float64 = 3.141592653589793238

	// PI1 = 4
	PI2 = 4

	fmt.Printf("원주율: %f\n", PI1)
	fmt.Printf("원주율: %f\n", PI2)
}
