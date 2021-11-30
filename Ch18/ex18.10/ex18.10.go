package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 삭제할 인덱스

	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}

	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
