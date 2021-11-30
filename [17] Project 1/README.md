# Chapter 17. 숫자 맞추기 게임
이번 챕터는 프로젝트 숫자 맞추기 게임을 개발하는 것이다. 

그래서 지금까지 배운 내용을 바탕으로, 책을 보고 따라하는 것이 아니라 직접 코드를 짜보기로 했다.

### 💡 숫자 맞추기 게임 방식

1. 콘솔에 예상 숫자 입력
2. 청답이면 프로젝트 종료
3. 더 작거나, 크면 해당 메시지 출력
4. 숫자를 맞출 때까지 반복

## 풀이

// 1. 현재 시간을 씨드로 넣어, 난수를 발생시킨다. (범위는 1~100)
// 2. 반복문 이용하여 Scanln 입력 받기를 반복
// 3. 그 안에서 if문 이용하여 정답이 맞으면 break, 틀리면 계속 반복

### code (1)

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. 현재 시간을 씨드로 넣어, 난수를 발생시킨다. (범위는 1~100)
	// 2. 반복문 이용하여 Scanln 입력 받기를 반복
	// 3. 그 안에서 if문 이용하여 정답이 맞으면 break, 틀리면 계속 반복

	fmt.Println("Start! Enter your number.")

	// 1
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(100) // 0~99
	var input int
	var cnt int

	// 2
	for { // for ans!=input 도 가능
		fmt.Scanln(&input)
		cnt++
		// 3
		if input > ans {
			fmt.Println("UP!")
		} else if input < ans {
			fmt.Println("DOWN!")
		} else {
			break
		}
	}
	fmt.Printf("Correct!\nYour tries: %d", cnt)
}
```
### -> 작동 X


이유는, input으로 여러 번 받는다고 해도, 처음에 받았던 값이 유지되기 때문이었다.

그래서 **버퍼를 비워줘야** 한다는 걸 깨달았다!

그리고 위와 같이 입력값으로 문자열이 들어왔을 때는, 에러메시지를 출력하도록 바꿔봐야겠다.

### code (2)

버퍼에 접근하기 위해 bufio 패키지와 os 패키지를 추가한다.

그리고 아래와 같이 InputIntValue() 함수를 넣어 입력값을 검증하고, 버퍼를 지워주는 작업을 한다.

```go
var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

```

main에서는 이렇게 사용한다.

```go
input, err := InputIntValue() // 숫자값 입력
		if err != nil {
			fmt.Println("Wrong Input. Enter integer only.")
		} else { // 여기에 대소 비교 코드 입력 }
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/42ad5bbd-17d1-466c-8dea-9967fa38abd0/Untitled.png)

### 최종 코드

```go
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
```
