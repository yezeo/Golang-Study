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
	fmt.Println("Start!")

	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(100) // 랜덤값 생성
	cnt := 0

	for {
		fmt.Printf("Guess Integer: ")
		input, err := InputIntValue() // 숫자값 입력
		if err != nil {
			fmt.Println("Wrong Input. Enter integer only.")
		} else {
			if input > ans { // 숫자값 비교
				fmt.Println("DOWN!")
			} else if input < ans {
				fmt.Println("UP!")
			} else {
				fmt.Println("Correct! Your tries: ", cnt)
				break
			}
			cnt++
		}
	}
}
