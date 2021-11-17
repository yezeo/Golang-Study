package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50} //배열 크기 생략

	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
