package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // 정렬되지 않은 슬라이스
	sort.Ints(s)                 // 정렬
	fmt.Println(s)
}
