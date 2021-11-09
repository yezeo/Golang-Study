# Chapter 11. for
## **11.1 for문 동작 원리:** `for 초기문; 조건문; 후처리 { }`

Go 언어는 반복문으로 for문 하나만 지원하지만, 여러 형태가 있기 때문에 각 형태를 적재적소에 잘 사용해야함.

### for문 동작 순서

초기문 먼저 실행    

→ 조건문 검사 

→ 조건문 결과가 true이면 for문 { 안쪽 코드 블록 } 수행 

→ 후처리 구문 → 다시 조건문 검사

만약 조건문이 false이면 후처리 없이 for문 종료

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // 초기문; 조건문; 후처리
		fmt.Print(i, ", ") //i 값 출력
	}

	//  fmt.Println(i)  // Error - i는 이미 사라졌다.
}`
```

→ `0, 1, 2, 3, 4, 5, 6, 7, 8, 9` 출력.

- i 변수 선언하고 0으로 초기화, i를 10과 비교, 조건문 true → {} 코드 블록 실행 → 후처리 i++ 실행
- 반복 후 i가 10이 되면 조건문 10<10은 false가 되어 for문 종료
- for문 안에서 선언된 i가 for문이 종료되면서 메모리에서 제거되기 떄문에 for문 밖에서 i를 호출하면 에러 발생

### 😮 **초기문 생략**

```go
for ; 조건문; 후처리 {
	코드 블록
}
```

### 😮 **후처리 생략**

```go
for 초기문; 조건문; {
	코드 블록
}
```

### 😮 **조건문만 있는 경우**

```go
for ; 조건문; {
	코드 블록
}
```

또는

```go
for 조건문 {
	코드 블록
}
```

### **11.1.4 무한 루프**

- 조건문이 true이면 코드 블록이 무한 반복되는 무한 루프가 됨.

```go
for true {
 코드 블록
}
```

switch문에서 조건문을 생략하면 true가 되듯이, for문에서도 **true 생략 가능**

```go
for {
 코드 블록
}
```

### 📌 주의.

무한 루프는 프로그램이 강제 종료되거나 break를 사용해 for문을 종료하지 않으면 계속 반복되니 사용에 주의를 기울여야 함.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}
```

→ `1 2 3 4 5 ...` 출력

- 위 예제는 프로그램을 강제 종료하지 않으면 계속 숫자 출력.
    
    → **Ctrl + C**를 눌러 강제 종료하기!
    

## **11.2 `continue`와 `break`**

continue와 break는 반복문을 제어하는 키워드.   

- `continue`: 이후 코드 블록을 수행하지 않고 곧바로 후처리
- `break`: for문에서 곧바로 빠져나옴

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // 무한 루프
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number) // 한줄 입력 받기
		if err != nil {               // 숫자가 아닌 경우
			fmt.Println("숫자를 입력하세요")

			// 키보드 버퍼를 비웁니다.
			stdin.ReadString('\n') // 키보드 버퍼 지우기
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다\n", number)
		if number%2 == 0 { // 짝수 검사
			break // for문 종료
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
```

- 입력받은 수가 짝수일 때까지 계속 입력을 받아서 출력하는 예제
- **줄바꿈까지 문자열을 다시 읽어서** 키보드 버퍼를 없애줘야함
- continue를 통해서 이후 코드 블록을 건너뜀.    
무한 루프라서 후처리와 조건검사가 없기 때문에 바로 다음 반복 시작

## **11.3 중첩 for문**

```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ { // 3번 반복
		for j := 0; j < 5; j++ { //5번 반복
			fmt.Print("*") // * 출력
		}
		fmt.Println()
	}
}
```

→ 결과 

```go
****
*****
*****
```

- 이중 중첩된 for문에서 각 for문이 i번, j번 반복하면 `i * j`번 반복하게 됨.
    
    삼중 중첩 for문은? 총 `i * j * k`번 반복.     
    
    → 중첩 반복문을 사용하면 연산량을 크게 증가시키므로, 반복 횟수가 많을 떄는 사용에 특히 주의해야 함.
