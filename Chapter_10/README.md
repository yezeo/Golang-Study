# Chapter 10. switch 문

## 10-I. switch문 동작 원리

- switch 키워드 다음 비굣값 옴
- 첫 번째 case부터 값을 검사
    - 만약 비굣값 == case값이면 case 문장 수행 후 **switch문 종료**
- 같은 값이 없으면 **deafult**문 시행
    - default는 생략 가능

---

## 10-II. switch문을 이용하는 경우?

복잡한 if else문을 보기 좋게 정리 가능

```go
if day == 1{
    fmt.Println("첫째 날입니다.")
    fmt.Println("오늘은 팀 미팅이 있습니다.")
}
...
if day == 5 {
    fmt.Println("다섯째 날입니다.")
    fmt.Println("최종 계약하는 날입니다.")
}
```

날짜별로 다른 메시지 출력하는 예제에 else if 너무 많아서 동작이 눈에 잘 보이지 않음

```go
switch day {
case 1:
    fmt.Println("첫째 날입니다.")
    fmt.Println("오늘은 팀 미팅이 있습니다.")
...
}
```

switch문을 이용하면 if문 사용한 예제보다 **가독성**이 좋음.

---

## 10-III. 다양한 switch문 형태

### 10-III-a. 한 번에 여러 값 비교

하나의 case는 하나 **이상의 값**을 비교할 수 있다. 각 값은 `쉼표`로 구분한다.

```go
// ex10.4
case "monday", "tuesday":
```

### 10-III-b. 조건문 비교

- **true**가 되는 조건문을 검사할 수 있다.

```go
switch true {
    case temp < 10, temp > 30:
```

switch문은 비굣값과 case의 값이 같아지는 경우를 찾는 구문이므로 **비굣값을 true로 할 경우 case의 조건문이 true가 되는 경우에 실행**됨

```go
switch true {
    ...
}
```

그런데 switch 다음에 비굣값을 적지 않는 경우 *default값으로 true*을 이용함

```go
switch {
    ...
}
```

그래서 위처럼 줄여 쓸 수 있음

### 10-III-c. switch 초기문

if문처럼 switch문도 초기문을 넣을 수 있음

```go
switch 초기문; 비굣값 {
case 값1:
    ...
case 값2:
    ...
default:
}
```

---

## 10-IV. const 열거값과 switch

`const 열거값`에 따라 수행되는 로직을 변경할 때 switch문을 이용함

- ex10.8: switch 문을 이용해 색깔을 나타내는 열거값을 문자로 바꿈

---

## 10-V. `break`와 `fallthrough` 키워드

- Go 언어: break를 사용하지 않아도 **case 하나 시행 후 자동으로 switch문 빠져나감**
    - 다른 언어는 case 종료 시에 break문 사용해야 함


- 하나의 case 문 시행 후 **다음 case문까지 같이 실행하고 싶을 때**: `fallthrough` 키워드 이용
    - 키워드 이용하면 다음 case까지 실행됨


```go
switch a {
    case 1:
        fmt.Println("a == 1")
        fallthrough
    case 2:
        fmt.Println("a == 2")
}

// a가 1일 때 a == 1 뿐 아닌 a == 2 구문이 같이 실행되어 출력됨
```

그런데 코드를 보는 사람에게 혼동을 일으킬 수 있으니 되도록 사용하지 않기를 권장함

---

## 연습 문제

1. switch, case, case, default
2. 긍정적인 평가
3. 
```go
package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(angle float64) Direction {
	switch true {
	case angle >= 315:
		return North
	case angle < 45 && angle >= 0:
		return North
	case angle >= 45 && angle <= 135:
		return East
	case angle > 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return North
	default:
		return None
	}
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
```