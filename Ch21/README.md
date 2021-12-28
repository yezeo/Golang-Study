# Chapter 21 함수 고급편

## 21.1 가변 인수 함수

가변 인수 함수 : 함수 인수 개수가 고정적이지 않은 함수

### 21.1.1 ... 키워드 사용

... 키워드를 사용해서 가변 인수 처리

```go
package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
```

출력문

```go
nums 타입: []int
15
nums 타입: []int
30
nums 타입: []int
0
```

- 인수 타입 int 앞에 ... → 가변 인수 타입
- sum() 함수 내부에서 nums는 int 슬라이스 타입 []int 처리

여러 타입을 인수로 섞어 쓰도록 구현 → 빈 인터페이스 interface{}

모든 타입이 빈 인터페이스 포함 → 빈 인터페이스 가변 인수 ...interface{} 타입

## 21.2 defer 지연 실행

파일 작업 이후 반드시 파일 핸들을 반환해야하기 때문에 함수 종료 전에 처리해야하는 코드가 있을 때 defer를 사용 → 해당 함수가 종료되기 직전 명령문 실행

```go
defer 명령문
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다")

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}
```

출력문

```go
파일에 Hello World를 씁니다.
파일을 닫았습니다
반드시 호출됩니다.
```

## 21.3 함수 타입 변수

함수 타입 변수 : 함수를 값으로 갖는 변수 (함수를 숫자로?)

프로그램 카운터 : 다음 실행할 라인을 나타내는 레지스터

함수 시작 지점 = 함수 포인터 → 숫자, 즉 변수의 값이 될 수 있다

함수 타입은 함수명과 함수 코드 블록을 제외한 함수 정의로 표시

```go
func add(a, b int) int {
	return a + b
}
```

```go
func (int, int) int
```

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
```

출력문

```go
12
```

- getOperator()의 반환 타임 : func (int, int) int → 함수 타입 정의
- getOperator(”*”) → mul() 함수 반환, operator값은 mul() 함수를 가리킨다

**별칭으로 함수 정의 줄여 쓰기**

```go
type opFunc func (int, int) int
func getOperator(op string) opFunc
```

**함수 정의에서 매개변수명 생략하기**

```go
func (int, int) int
func (a int, b int) int
```

## 21.4 함수 리터럴

함수 리터럴 : 이름 없는 함수 (함수명 X, 함수 타입 변숫값으로 대입되는 함숫값)

```go
package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {

		return func(a, b int) int { // 함수 리터럴
			return a + b
		}
	} else if op == "*" {

		return func(a, b int) int { // 함수 리터럴
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4) 
	fmt.Println(result)
}
```

출력문

```go
12
```

### 21.4.1 함수 리터럴 내부 상태

함수 리터럴은 필요한 변수를 내부 상태로 가질 수 있다

```go
package main

import "fmt"

func main() {
	i := 0

	f := func() {
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}
```

출력문

```go
11
```

- i 변수는 함수 내부에 있지 않는 외부 변수
- 함수 리터럴 내부에서 외부 변수 i에 접근 → 내부 상태로 가져와 접근

### 21.4.2 함수 리터럴 내부 상태 주의점

캡쳐 : 함수 리터럴 외부 변수 → 내부 상태로 가져오는 것

캡쳐는 값 복사가 아닌 참조 형태

```go
package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
```

출력문

```go
CaptureLoop
3
3
3
CaptureLoop2
0
1
2
```

- f := make([]func() , 3) → f는 func() 타입 함수 리터럴 3개 갖는 슬라이스
- 3, 3, 3이 호출된 이유 : i 변수 캡쳐할 때 참조로 캡쳐되기 때문
- → i값을 저장하는 변수를 만들어 캡쳐하면 0, 1, 2 출력

**참조로 변수 가져오기?** 변수의 주소를 포인터 값으로 복사

### 21.4.3 파일 핸들을 내부 상태로 사용하는 예

함수 리터럴을 이용해서 원하는 함수를 그때그때 정의해서 함수 타입 변숫값으로 사용, 필요한 외부 변수를 내부 상태로 가져와서 편리하게 사용

```go
package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
```

- 파일에 msg를 쓰는 함수 리터럴 → writeHello() 함수의 인수로 사용
- writeHello() 함수 내부에서 정의한 함수 리터럴의 인수(msg): “Hello World”

**의존성 주입** : 외부에서 로직을 주입하는 것

## 연습문제

1. ...
2. Result: 30