# Chapter 10. Switch
## 10.1 **switch문 동작 원리**

switch문은 **값**에 따라 다른 로직을 수행할 때 사용.

```go
a := 3

switch a { 
case 1:
	fmt.Println("a == 1")
case 2:
	fmt.Println("a == 2")
case 3: 
	fmt.Println("a == 3")
case 4:
	fmt.Println("a == 4")
default:
	fmt.Println("a > 4")
}
```

→ `a == 3` 출력됨.

- 비굣값과 case값이 같으면 해당 case 문장 수행 후 switch문을 종료한다.
- 같은 값이 없으면 default 문장 수행 (default는 생략 가능)

## 10.2 **switch문을 언제 쓰는가?**

switch문을 이용하여, 복잡한 if-else문을 보기 좋게 정리할 수 있다.   

비교해보자.

### if문을 사용할 때

```go
package main

import "fmt"

func main() {

	day := 3

	if day == 1 {
		fmt.Println("첫째 날입니다")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	} else if day == 2 {
		fmt.Println("둘째 날입니다")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	} else if day == 3 {
		fmt.Println("셋째 날입니다.")
		fmt.Println("설계안을 확정하는 날입니다.")
	} else if day == 4 {
		fmt.Println("넷째 날입니다.")
		fmt.Println("예산을 확정하는 날입니다.")
	} else if day == 5 {
		fmt.Println("다섯째 날입니다.")
		fmt.Println("최종 계약하는 날입니다.")
	} else {
		fmt.Println("프로젝트를 진행하세요")
	}
}
```

### switch문을 사용할 때

```go
package main

import "fmt"

func main() {

	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날입니다")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날입니다")
		fmt.Println("오늘은 면접이 있습니다.")
	case 3:
		fmt.Println("셋째 날입니다.")
		fmt.Println("설계안을 확정하는 날입니다.")
	case 4:
		fmt.Println("넷째 날입니다.")
		fmt.Println("예산을 확정하는 날입니다.")
	case 5:
		fmt.Println("다섯째 날입니다.")
		fmt.Println("최종 계약하는 날입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요")
	}
}
```

if문을 사용한 예제보다 switch문을 사용한 예제가 더 가독성이 좋다.

→ 출력

```go
셋째 날입니다.
설계안을 확정하는 날입니다.
```

## **10.3 다양한 switch문 형태**

### **10.3.1 한 번에 여러 값 비교** (쉼표 , 사용)

- 하나의 case가 하나 이상의 값을 비교할 수 있다.

```go
package main

import "fmt"

func main() {

	day := "thursday"

	switch day {
	case "monday", "tuesday": 
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}
```

→ `수, 목, 금요일은 실습 가는 날입니다.` 출력

### **10.3.2 조건문 비교**

- 단순히 값만 비교하는 게 아니라, if문처럼 true가 되는 조건문을 검사할 수 있음

```go
package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10 || temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요")

	// 이미 두 번째 case를 실행했으므로 아래는 검사 X
	case temp >= 15 && temp < 25:
		fmt.Println("야외활동 하기 좋은 날씨입니다")
	default:
		fmt.Println("따뜻합니다.")
	}
}
```

→ `약간 추울 수 있으니 가벼운 겉옷을 준비하세요`  출력.

### 📌

비굣값을 true로 할 경우, **case의 조건문이 true가 되는 경우**가 실행된다.

- switch문 다음에 비굣값을 적지 않은 경우 default값으로 true 사용
    
    ```go
    switch {
    	... // 조건문이 true가 되는 경우를 실행한다.
    }
    ```
    

### **10.3.3** `switch 초기문; 비굣값`

- if문과 마찬가지로 switch문에서도 초기문 사용 가능

```go
package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); age { // getMyAge() 결과값 비교
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age) // age값 사용
	}

	fmt.Println("age is", age) // error - age 변수는 사라짐
}
```

→ `.\ex10.6.go:19:24: undefined: age` 출력.

- 초기문으로 getMyAge() 함수 실행하여 age 반환, 비굣값은 age값
- age는 **switch문이 종료되기 전까지만** 접근 가능
    
    → switch문 종료되면 age도 사라지기 때문에 에러 발생
    

```go
package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	// age 변수 선언 및 초기화
	switch age := getMyAge(); true { 
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age) // age값 사용
	}
}
```

출력문

`20s`

- age 변수를 초기문을 통해서 선언 및 대입 후 switch의 비굣값으로 true 사용. true일 경우 생략 가능하므로 다음과 같이 작성 가능

`switch age := getMyAge(); {`

## **10.4 const 열거값과 switch**

const 열거값에 따라 수행되는 로직을 변경할 때 switch문을 주로 사용.

```go
package main

import "fmt"

//ColorType 별칭 타입을 정의
type ColorType int
const (
	Red ColorType = iota // iota로 열거값 정의 (0부터 시작해 1씩 증가)
	Blue
	Green
	Yellow
)

// 각 ColorType 열거값에 따라 문자열 반환하는 함수
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}
func getMyFavoriteColor() ColorType {
	return Green
}
func main() {
	fmt.Println("My Favorite color is", colorToString(getMyFavoriteColor()))
}
```

→ `My favorite color is Blue` 출력   

- ColorType 별칭 타입을 정의하고 열거값을 정의. switch문을 이용해 각 열거값에 따른 문자열을 반환하는 함수.
- 위 예에서 새로운 색깔을 추가한다면 colorToString() 함수도 수정해야함.

이런 경우 "커플링되었다", "결함되어있다"라고 한다. 열거값이 수정될 때 연관된 모든 switch case문도 수정해야함. 그래서 열거값에 연관된 switch case가 많아질수록 작은 수정에도 많은 코드가 변경되어야하는 **산탄총 수술 문제**가 발생. **하나의 열거값에 연관된 switch case는 최대한 줄이는 게 좋음.**

## **10.5 `break`와 `fallthrough` 키워드**

Go 언어에서는 `break`를 사용하지 않아도 case 하나를 실행 후 자동으로 switch문을 빠져나간다. (다른 언어와의 차이점! )

```go
package main

import "fmt"

func main() {

	a := 3

	switch a { 
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3: 
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}
```

→ `a == 3` 출력

- break를 사용하든 않든 상관없이 case 하나를 실행 후 switch문을 빠져나감.
- 만약 하나의 case문 실행 후 다음 case문까지 같이 실행하고 싶을 때 → `fallthrough` 키워드 사용

```go
package main

import "fmt"

func main() {

	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	default:
		fmt.Println("a > 4")
	}
}
```

→ 출력문

```go
a == 3
a == 4
```

# Chapter 10. **연습문제**

1. [코드](prob1/prob1.go)
2. 긍정적인 평가
3. [코드](prob3/prob3.go)
