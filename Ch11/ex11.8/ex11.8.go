package main

import "fmt"

func main() {
	a := 1
	b := 1

OuterFor: // 레이블 정의
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor // 레이블에 가장 먼저 포함된 for문까지 종료
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
