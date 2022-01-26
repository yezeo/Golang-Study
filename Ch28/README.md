# Chapter 28 테스트와 벤치마크

## 28.1 테스트 코드

Go 언어는 테스트 코드 작성과 실행을 언어 자체에서 지원

3가지 표현 규약을 따라 테스트 코드 작성, go test 명령으로 실행

1. 파일명이 _test.go로 끝나야 한다
2. testing 패키지를 임포트해야 한다
3. 테스트 코드는 func TestXxxx(t *testing.T) 형태이어야 한다

### 28.1.1 테스트 코드 작성하기

테스트 대상 코드 : 81을 반환하는 함수

```go
package main

import "fmt"

func square(x int) int {
	return 81
}

func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))
}
```

출력문

```go
9 * 9 = 81
```

테스트 코드 : 9의 제곱값이 81임을 테스트하는 코드

```go
package main

import (
	"testing"
)

func TestSquare1(t *testing.T) {

	rst := square(9)
	if rst != 81 {
		t.Errorf("squaure(9) should be 81 but square(9) returns %d", rst)
	}
}
```

- t.Errorf() 메서드에 테스트 실패 시 실패를 알리고 실패 메시지를 넣을 수 있다
- testing.T 객체의 Error() 메서드 : 테스트 실패 시 모든 테스트 중단
- testing.T 객체의 Fail() 메서드 : 테스트 실패해도 다른 테스트 진행

```go
func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("squaure(9) should be 81 but square(3) returns %d", rst)
	}
}
```

- 실패 : equare() 함수를 고쳐야 함

```go
func square(x int) int {
	return x * x
}
```

### 28.1.2 일부 테스트만 실행하기

모든 테스트를 실행하지 않고 특정 테스트만 실행

```go
go test -run 테스트명
```

### 28.1.3 테스트를 돕는 외부 패키지

“stretchr/testify” 패키지 : 테스트 코드 작성을 돕는 외부 패키지, 다양한 함수 제공

```go
go get github.com/stretchr/testify
```

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)                               //테스트 객체 생성
	assert.Equal(81, square(9), "square(9) should be 81") //테스트 함수 호출
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
```

- asset 객체 : 테스트 코드를 쉽게 만들 수 있는 다양한 메서드 포함 (Equal() 등)

**stretchr/testify/asset 패키지에서 제공하는 유용한 함수**

- Equal() : expected와 actual 두 값을 비교하여 다를 경우 테스트 실패 & 메시지 출력
- Greater() : e1이 e2보다 크지 않으면 테스트 실패 & 메시지 출력
- Len() : object의 항목 개수가 length가 아니면 테스트 실패 & 메시지 출력
- NotNilf() : object가 nil이면 테스트 실패 & 메시지 출력
- NotEqual() : expected와 actual이 같으면 테스트 실패 & 메시지 출력

**stretchr/testify 패키지에서 제공하는 그외 유용한 기능**

- mock 패키지 : 모듈의 행동을 가장하는 목업 객체 제공
- suite 패키지 : 테스트 준비 작업이나 테슽 종료 후 뒤처리 작업을 쉽게 할 수 있도록 도와주는 패키지

## 28.2 테스트 주도 개발

테스트의 중요성

1. 과거에 비해서 프로그램의 규모가 커졌다. : 많은 프로그래머의 코드들의 접점이 늘어 예기치 못한 버그도 늘 수 밖에 없다.
2. 과거에 비해 고가용성에 대한 요구사항이 높아졌다. 무중둔 서비스에 대한 사용자 요구가 늘어나면서 치명적인 버그를 줄이는 게 큰 역할이 되었다.

**블랙박스 테스트**

- 제품 내부를 오픈하지 않은 상태에서 진행되는 테스트
- 사용성 테스트
- 프로그램 실행한 상태로 실행 동작을 검사하는 방식

**화이트박스 테스트**

- 프로그램 내부 코드를 직접 검증하는 방식
- 유닛 테스트
- 전통적인 방식 : 코드 작성 → 테스트하고 버그를 발경 → 코드 수정 (많은 문제점을 가짐)
    - 빈약한 테스트 케이스 : 예외 상황이나 경계 체크가 무시되기 쉽다
    - 테스트 통과를 목적으로 하는 형식적인 테스트 코드를 작성하기 십상이다

**테스트 주도 개발**

- 방식 : 테스트 코드 작성 → 테스트 실패 → 코드 작성 → 테스트 성공 → 개선
- 개선은 SOLID 원칙에 입각해 진행 (= 리팩터링)
- 이점
    - 테스트 코드가 자연적으로 촘촘해진다
    - 개발이 재밌다 : ‘작은 목표 설정 → 실패 → 달성 → 달성 강화 → 새로운 작은 목표 설정’ 절차

## 28.3 벤치마크

Go 언어는 테스트 외 코드 성능을 검사하는 벤치마크 기능도 지원한다.

표현 규약

1. 파일명이 _test.go로 끝나야 한다.
2. testing 패키지를 임포트해야 한다.
3. 벤치마크 코드는 func BenchmarkXxxx(b *testing.B) 형태이어야 한다.

테스트 대상 코드 : 피보나치 수열 값을 구하는 두 가지 방식의 함수

```go
package main

import "fmt"

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2) //재귀 호출
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ { //반복문
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
```

테스트 코드 : 피보나치 함수 테스트

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(2) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(2) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ { 
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
```

- 벤치마크테스트 실행 명령

```go
go test -bench .
```

```go
goos: windows
goarch: amd64
pkg: musthaveGo/ex28.2
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkFibonacci1-4               7824            441022 ns/op
BenchmarkFibonacci2-4           38136084                29.85 ns/op
PASS
ok      musthaveGo/ex28.2       6.342s
```

- 벤치마크 기능을 이용해서 손쉽게 코드 성능을 측정하고 비교해서 더 나은 성능을 가진 알고리즘으로 교체 가능

## 연습문제

1. 1 (파일명은 _test.go로 끝나야 한다) 2 (주요 기능 뿐만 아니라 모든 코드의 모든 기능을 검사해야 한다) 3 (테스트 함수명은 TestXxx로 시작해야 한다) 4 (벤치마크 함수 형식은 func BenchmarkXxx(b *testing.B) 형태이다)
2. 테스트 주도 개발
3.  

```go
package main

import (
	"testing"
)

func TestAtoi1(t *testing.T) {
	n, err := Atoi("0")
	if err != nil {
		t.Fail()
	}
	if n != 0 {
		t.Fail()
	}
}

func TestAtoi2(t *testing.T) {
	n, err := Atoi("1")
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}

	n, err = Atoi("5")
	if err != nil {
		t.Fail()
	}
	if n != 5 {
		t.Fail()
	}
}

func TestAtoi3(t *testing.T) {
	n, err := Atoi("12")
	if err != nil {
		t.Fail()
	}
	if n != 12 {
		t.Fail()
	}

	n, err = Atoi("3523")
	if err != nil {
		t.Fail()
	}
	if n != 3523 {
		t.Fail()
	}
}

func TestAtoi4(t *testing.T) {
	n, err := Atoi("ab34c")
	if err == nil {
		t.Fail()
	}

	n, err = Atoi("  3523")
	if err != nil {
		t.Fail()
	}
	if n != 3523 {
		t.Fail()
	}
}

func TestAtoi5(t *testing.T) {
	n, err := Atoi("-12")
	if err != nil {
		t.Fail()
	}
	if n != -12 {
		t.Fail()
	}
}
```