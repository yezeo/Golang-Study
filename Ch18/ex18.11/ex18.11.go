package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0) // 맨 뒤에 요소 추가

	idx := 2 // 추가하려는 위치

	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	slice[idx] = 100

	fmt.Println(slice)
}
