# Chapter 11. for문

## 11-I. for문 동작 원리

- Go 언어에서는 반복문으로 for문 하나만 지원

```go
for 초기문; 조건문; 후처리 {
    코드 블록
}
```

### 11-I-a. 초기문 생략

```go
for ; 조건문; 후처리 {
    코드 블록
}
```

초기문을 생략해도 ;를 붙여서 조건문 자리 표시해 주어야 함

### 11-I-b. 후처리 생략

```go
for 초기문; 조건문; {
    코드 블록
}
```

후처리를 생략해도 조건문 뒤에 ; 붙여 줘야 함

### 11-I-c. 조건문만 있는 경우

```go
for ; 조건문; {
    코드 블록
}
```

혹은 더 단순하게 ; 없앨 수도 있음

```go
for 조건문 {
    코드 블록
}
```

### 11-I-d. 무한 루프

**조건문이 true**이면 코드 블록이 무한 반복되는 **무한 루프**가 됨

```go
for true {
    코드 블록
}
```

switch문처럼 조건문에서 true 생략 가능

```go
for {
    코드 블록
}
```

무한 루프는 프로그램이 종료되거나 break를 사용해 종료하지 않으면 계속 반복됨

```go
// 1초마다 한 번씩 1씩 증가하는 숫자를 무한히 출력하는 예문
package main

import (
    "fmt"
    "time"
)

func main() {
    time.Sleep(time.Second)
    for {
        fmt.Println(i)
        i++
    }
}
```

---

## II-2. continue와 break

- 반복문을 제어하는 키워드
    - `continue`: 이후 코드 블록 실행 X. **바로 후처리 후 조건문 검사**
    - `break`: for문에서 곧바로 빠져나옴

---

## II-4. 중첩 for문과 break, 레이블

- break 이용: break가 속한 for문에서만 빠져나옴
- **모든 for문**을 빠져나가고 싶을 때는 `불리언 변수`를 이용함

```go
found := true
if found {
    break
}
```

- `레이블`을 이용하는 방법
    - 되도록이면 플래그 사용, 레이블은 꼭 필요한 경우에만 사용하기를 권장

```go
// 레이블을 사용해서 for문을 종료하는 코드
Outerfor: // 레이블 정의
    for {
        for {
             if {
                break OuterFor // 레이블에 가장 먼저 포함된 for문까지 종료
            }   
        }
    }
```

- 클린 코드를 지향하려면 중첩된 내부 로직을 **함수로 묶어** 중첩을 줄이고, 플래그 변수나 레이블 사용을 **최소화**해야 함

---

## 연습 문제

1. 
```go
package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {
		fmt.Print(i, " ")
	}
}
```

2. 
```go
		if i >= 3 && i <= 6 {
			continue
		}
```

3.
```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i, "*", i, "=", i*i)
		}
	}
}
```

4. 
```go
package main

import "fmt"

func main() {
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
```