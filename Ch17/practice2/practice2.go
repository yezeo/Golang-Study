package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	balance := 1000

	for {
		r := rand.Intn(5) + 1
		fmt.Print("1~5사이의 값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5사이의 값만 입력하세요.")
		} else {
			if n == r {
				balance += 500
				fmt.Println("맞췄습니다! 남은 돈: ", balance)
				if balance >= 5000 {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= 100
				fmt.Println("틀렸습니다! 남은 돈: ", balance)
				if balance <= 0 {
					fmt.Println("게임 오버")
					break
				}
			}
		}
	}
}
