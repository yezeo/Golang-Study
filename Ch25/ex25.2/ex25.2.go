package main

import "fmt"

func main() {
	ch := make(chan int) //크기 0인 채널

	ch <- 9
	fmt.Println("Never print") //실행되지 않는다
}
