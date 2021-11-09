# Chapter 09. if문


## 09-II. 그리고 &&, 또는 ||

### 09-II-a. 쇼트서킷

- **`쇼트서킷`**(short-circuit)
    - && 연산은 좌변이 false이면 **우변을 검사하지 않고** false 처리
    - || 연산 또한 좌변이 true이면 **우변을 검사하지 않고** true 처리
    - **조건문 우변이 실행되지 않을 수 있으므로** 염두에 두고 코드를 구현해야 함

- ex9.4: 쇼트서킷에 의해 IncreaseAndReturn() 함수가 호출되지 않음

### 09-II-b. 소괄호() 활용

- `소괄호`를 사용하여 더욱 다양한 조건문 생성 가능

```go
// 예약을 했거나, 가진 돈이 200이 넘고 빈자리가 있으면 통과
if (hasBooked == true) || (money > 200 && hasEmptySeat == true)
```

---

## 09-IV. if 초기문; 조건문

- if 조건을 검사하기 **전**에 **초기문**을 넣을 수 있다.
- 초기문은 **검사에 사용할 변수를 초기화할 때** 주로 사용한다.

```go
if 초기문; 조건문 {
    문장
}
```

초기문 자리에 **하나의 구문**이 올 수 있고, 끝에 `;`을 붙여 끝났음을 표시

```go
if filename, success := UploadFile(); success {
    fmt.Println("Upload success", filename)
} else {
    fmt.Println("Failed to upload")
}
```

- 먼저 UploadFile() 함수 실행, filename과 success에 변수 반환값을 저장
- 함수 성공 여부에 따라 다른 메시지 출력

if 초기문은 이렇게 **함수의 결과를 검사할 때 주로 사용**

초기문에서 선언한 변수의 범위는 if문 **안**으로 한정

---

## 연습 문제

1. if, else if, else
2. 신중히 생각해보세요.
3. 
```go
package main

import "fmt"

func main() {
    temp := 30
    rain := 40

    if temp >= 25 {
        if rain >= 80 {
            fmt.Println("덥고 비가 옵니다.");
        } else if rain >= 20 {
            fmt.Println("덥고 습합니다.");
        } else {
            fmt.Println("야외 활동 하기 좋습니다.");
        }
    } else if temp < 10 || rain >= 80 {
            fmt.Println("야외 활동 하기 좋지 않습니다.")
    } else {
        fmt.Println("좋은 날씨입니다.")
    }
}
```