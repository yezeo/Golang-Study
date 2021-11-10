package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // ❶ 무한 루프
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요")

			// 키보드 버퍼를 비웁니다.
			stdin.ReadString('\n') // 키보드 버퍼를 지워준다.
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다\n", number)
		if number%2 == 0 { // 짝수 검사
			break // for문 종료
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
