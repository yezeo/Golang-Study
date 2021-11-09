# 9장. if문
## if문 기본 사용법

`if` - `else if` - `else` 로 조건에 따라 분기할 수 있다.    

```go
package main

import "fmt"

func main() {
	light := "red"

	if light == "green" {
		fmt.Println("길을 건넌다")
	} else {
		fmt.Println("대기한다")
	}

}
```

light의 값이 "green"이 아니기 때문에 else 문장인 "대기한다"가 출력된다.   

else if와 else 구문은 생략이 가능하다. 

## 그리고 &&, 또는 ||

- AND 논리연산자 &&
- OR 논리연산자 ||

```go
package main

import "fmt"

func main() {
	var age = 22

	if age >= 10 && age <= 15 {
		//age가 10 이상 15 이하인 경우
		fmt.Println("You are Young")
	} else if age > 30 || age <20 {
		// age가 20대가 아닌 경우 (30보다 크거나, 20보다 작은 경우)
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best age of your life")
	}
}
```

### '쇼트서킷'

&& 연산은 좌변이 false이면 우변을 검사하지 않고 false 처리를 한다.   

|| 연산 역시 좌변이 true이면 우변을 검사할 필요가 없고 바로 true 처리를 한다.   

이를 **쇼트서킷** 이라고 한다.   

쇼트서킷이 실행됨으로써 조건문 우변이 실행되지 않는다는 점을 유의해야 한다. 아래 코드를 보자.

```go
package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 { // 함수가 호출되지 않는다.
		fmt.Println("1 increase")
	}
	if true || IncreaseAndReturn() < 5 { // 함수가 호출되지 않는다.
		fmt.Println("2 increase")
	}

	fmt.Println("cnt:", cnt)
}
```

결국 `IncreaseAndReturn()` 함수는 한 번도 호출되지 않았기 때문에 cnt 값은 증가되지 않고 0에서 변화되지 않았다.

## 중첩 if

복잡한 경우를 표현할 때, if문 안에 if문을 중첩하여 사용할 수 있다.

```go
package main

import "fmt"

// 친구 중 부자가 있는가 반환 - 무조건 true 반환
func HasRichFriend() bool {
	return true
}

// 같이 간 친구 숫자를 반환 - 무조건 3을 반환
func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 { 
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 { 
		if GetFriendsCount() > 3 { 
			fmt.Println("어이쿠 신발끈이..")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}
}
```

## `if 초기문; 조건문`

if문 조건을 검사하기 전에 초기문을 넣을 수 있다.   

초기문은 검사에 사용할 변수를 초기화할 때 사용한다.   

```go
if filename, success := UploadFile(); success {
	fmt.Println("Upload success", filename)
} else {
	fmt.Println("Failed to upload")
}
```

**주의)** 초기문에서 선언한 변수의 범위는 **if문 안으로 한정**된다.

```go
package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age) // age 값에 접근가능하다 
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	fmt.Println("Your age is", age) // Error - age는 소멸됐다
}
```

---

# ch9 연습문제

1. if, else if, else
2. "신중히 생각해보세요"가 출력된다.
3. [코드](./pr03.go)
