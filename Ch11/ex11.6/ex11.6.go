package main

import "fmt"

func main() {
	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 { //
				break // 안쪽 for문 종료
			}
		}
		b = 1
		dan++
		if dan == 10 {
			break // 바깥쪽 for문 종료
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
