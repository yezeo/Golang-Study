package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are Young", age)
	} else if ok && (age > 30 || age < 20) {
		fmt.Println("You are not 20s", age)
	} else if ok {
		fmt.Println("You age Beautiful", age)
	} else {
		fmt.Println("Error")
	}
	// fmt.Println("Your age is", age) 에러 발생, age 소멸됨
}
