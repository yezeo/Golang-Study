package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Print("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Print("히터를 켠다")
	} else {
		fmt.Println("대기한다")
	}
}
