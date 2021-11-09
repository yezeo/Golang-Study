package main

import "fmt"

func main() {
	var age = 22

	if age >= 10 && age <= 15 {
		//age가 10 이상 15 이하인 경우
		fmt.Println("You are Young")
	} else if age > 30 || age <20 {
		// age가 20대가 아닌 경우 (30보다 크거나, 20보다 작은 경우)
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best age of your life")
	}
}