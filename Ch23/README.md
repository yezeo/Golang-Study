# Chapter 23 에러 핸들링

## 23.1 에러 반환

가장 기본 방식: 에러 반환하고 알맞게 처리

예) ReadFile() 함수 → 해당 파일 없어 에러 발생 : 메세지 출력 후 다른 파일 읽거나 임시 파일 생성

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) 
	if err != nil {
		return "", err
	}
	defer file.Close()         
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) 
	if err != nil {                  
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) 
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) 
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") 
		if err != nil {                                
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename) 
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line)
}
```

출력문

```go
파일내용: This is WriteFile
```

ReadFile() 함수

- os.Open() 함수로 파일 열기 → 두 번째 반환값 error 실패 : 에러 발생, file.Close() 함수 defer 키워드 사용하여 파일 핸들 닫기
- bufio.NewReader() 함수로 객체 생성 : ‘\n’ 나올 때까지 파일에서 문자열 읽기, 두번째 반환값 error는 문자열이 delim 문자로 끝나지 않을 경우만 (밑줄 _ 로 에러 무시)

WriteFile() 함수

- os.Create() 함수로 파일 생성
- fmt.Fprintln() 함수로 파일 핸들에 문자열 작성: 첫번째 인수로 Write() 메서드를 가진 io.Writer 인터페이스, 두번째 인수로 쓸 내용, 첫번째 반환값으로 쓴 길이 (밑줄 _로 무시), 두번째 반환값으로 에러 발생 시 에러

main() 함수

- ReadFile() 함수 시도 → 에러 발생 시 WriteFile() 함수 호출로 파일 생성 → 에러 발생 시 에러 메세지 출력 후 종료 / 성공 시 ReadFile() 함수 시도 후 내용 출력

### 23.1.1 사용자 에러 반환

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(
			"제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
```

출력문

```go
Error: 제곱근은 양수여야 합니다. f:-2
```

- Sqrt() 함수 : 인수의 제곱근 반환, 인수 f가 음수이면 에러 반환
- fmt.Errorf() 함수로 에러 반환 → 원하는 에러 메시지 생성 가능
- errors 패키지의 New() 함수로 error 생성 가능

## 23.2 에러 타입

### 23.2.1 에러 랩핑

## 23.3 패닉

패닉 : 프로그램을 정상 진행시키기 어려운 상황에서 프로그램 흐름을 중지시키는 기능

내장 함수 panic() : 문제 발생 시점에 프로그램 바로 종료, 빠르게 문제 발생 시점 파악 가능 (버그 수정에 유용), 콜 스택 표시 (함수 호출 순서)

```go
package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
```

출력문

```go
9 / 3 = 3
panic: b는 0일 수 없습니다

goroutine 1 [running]:
main.divide(0x9, 0x3)
        C:/Users/WINDOWS10/Desktop/Golang-Study/Ch23/ex23.5/ex23.5.go:7 +0x10b
main.main()
        C:/Users/WINDOWS10/Desktop/Golang-Study/Ch23/ex23.5/ex23.5.go:14 +0x31
```

- divide() 함수의 제수 b가 0이면 panic() 호출, panic() 함수의 인수로 에러 발생 원인 → 문제 발생 이유와 위치 빠르게 파악 가능
- 콜 스택 : panic이 발생한 마지막 함수 위치부터 역순으로 호출 순서 표시

### 23.3.1 패닉 생성

```go
func panic(interface{})
```

```go
panic(42)
panic("unreachable")
panic(fmt.Errorf("This is error num: %d", num)
panic(SomeType{SomeData})
```

일반적으로 string 타입 메시지나 fmt.Errorf() 함수 이용

### 23.3.2 패닉 전파 그리고 복구

프로그램 배포 이후 복구할 수 있는 패닉이라면 복구 시도

panic 호출 순서를 거슬러 올라가며 전파, 어느 단계에서든 패닉은 복구된 시점부터 프로그램이 계속된다 → recover() 함수로 패닉 복구 가능, 호출 시점에 패닉 전파 중이면 panic 객체 반환 or nil 반환

```go
package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0)) 
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨") 
}
```

출력문

```go
f() 함수 시작
9 / 3 = 3
panic 복구 - 제수는 0일 수 없습니다.
프로그램이 계속 실행됨
```

- main() → f() → g() → h() 패닉 발생 → g() → f() → main() 으로 전파
- 패닉이 f()까지 전파됐으나 defer 사용 → recover() 패닉 복구 시도 : panic 메시지 출력
- recover()는 제한적으로 사용하는 게 좋다 : 복구하더라도 불안정할 수 있다 (복구되면 내부 상태를 깨끗하게 정리해서 다른 오류 발생 방지)

### 23.3.3 recover() 결과

내장 함수 recover()는 발생한 panic 객체 반환

```go
func recover() interface{}
```

recover()로 반환한 타입을 실제로 사용하려면 타입 검사 필요

```go
if r, ok := recover().(net.Error); ok {
	fmt.Println("r is net.Error Type")
}
```

## 연습문제

1. *os.PathError 타입
2. 숫자만 입력하세요. 문자: c

3.

```go
15
nums should be not empty
end of main
```