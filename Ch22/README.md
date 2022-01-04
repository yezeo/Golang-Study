# Chapter 22 자료구조

## 22.1 리스트

리스트는 기본 자료구조로 여러 데이터 보관

배열과 가장 큰 차이점: 배열은 연속된 메모리, 리스트는 불연속된 메모리에 데이터 저장

### 22.1.1 포인터로 연결된 요소

리스트는 각 데이터를 담고 있는 요소들을 포인터로 연결한 자료구조 (링크드 리스트)

```go
type Element struct {  //구조체
	Value interface{ }   //데이터 저장 필드 
	Next  *Element       //다음 요소 주소 저장 필드
	Prev  *Element       //이전 요소 주소 저장 필드
}
```

- Value: interface{ } 타입으로 어떤 타입값도 저장 가능
- Next, Prev: 인스턴스의 메모리 주소 → 일관성이 없는 불연속 자료구조

### 22.1.2 리스트 기본 사용법

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()       
	e4 := v.PushBack(4)   
	e1 := v.PushFront(1)  
	v.InsertBefore(3, e4) 
	v.InsertAfter(2, e1) 

	for e := v.Front(); e != nil; e = e.Next() { 
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { 
		fmt.Print(e.Value, " ")
	}
}
```

출력문

```go
1 2 3 4
4 3 2 1
```

- list.New() 함수: 새로운 리스트 인스턴스 생성
- PushBack(), PushFront() 메서드: 리스트 맨 뒤/앞에 요소 추가, 요소 반환
- InsertBefore(), InsertAfter() 메서드: 두번째 인수로 입력된 요소 앞/뒤에 요소 추가
- Front(), Back() 메서드: 가장 첫 번째, 마지막 요소 반환
- Next(), Prev() 메서드: 현재 요소의 다음, 이전 요소 반환

### 22.1.3 배열 vs 리스트

차이점 이해하기

**맨 앞에 데이터 추가하기**

배열: 각 요소를 한 칸씩 뒤로 밀어냄, 맨 앞의 값을 변경 **→ O(N)**

리스트: 맨 앞에 요소를 추가하고 연결 만들기 **→ O(1)**

**특정 요소에 접근하기**

배열에서 인덱스 이동 공식 → 배열 시작 주소 + (인덱스 x 타입 크기) **→ O(1)**

리스트에서 앞 요소들을 모두 거쳐서 접근 **→ O(N)**

| 행위 | 배열, 슬라이스 | 리스트 |
| --- | --- | --- |
| 요소 삽입 | O(N) | O(1) |
| 요소 삭제 | O(N) | O(1) |
| 인덱스 요소 접근 | O(1) | O(N) |

**데이터 지역성** : 필요한 데이터가 인접할수록 처리 속도 빨라진다 → 배열이 리스트보다 데이터 지역성이 월등하게 좋다.

### 22.1.4 실습 : 큐 구현하기

큐 특징

- 들어간 순서 그대로 빠져나오기 때문에 순서 유지
- 새로운 요소는 항상 맨 마지막에 추가
- 출력값은 맨 앞에서 하나씩 뺌

리스트로 구현: O(1) 보장

```go
package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
```

출력문

```go
1 -> 2 -> 3 -> 4 ->
```

- 리스트 이용하여 큐 구조체 정의
- Push() 메서드 → PushBack() 메서드 활용
- Pop() 메서드 → Front() 메서드 + Remove() 메서드 활용

### 22.1.5 실습 : 스택 구현하기

스택 특징

- 가장 최근에 넣은 것부터 역순으로 나옴
- 요소는 맨 뒤로 추가
- 맨 뒤에서 뺌

```go
package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
```

출력문

```go
4 -> 3 -> 2 -> 1 ->
```

- Pop() 메서드 → Back() 메서드 활용
- 보통 큐는 리스트, 스택은 배열로 구현

## 22.2 링

링: 맨 뒤의 요소와 맨 앞의 요소가 서로 연결된 자료구조, 리스트 기반, 환형 리스트

```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}
```

출력문

```go
A B C D E
A E D C B
```

- 요소가 5개인 링 생성, 첫 번째 요소 인스턴스 반환
- 각 요소 순회하고 현재 위치 r은 다시 첫 번째 요소로

### 22.2 링은 언제 쓸까?

저장할 개수가 고정되고, 오래된 요소는 지워도 되는 경우

- 실행 취소 기능
- 고정 크기 버퍼 기능
- 리플레이 기능

## 22.3 맵

키와 값 형태로 데이터를 저장하는 자료구조

리스트나 링과 달리 container 패키지가 아닌 Go 기본 내장 타입

```go
package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구"

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"])
	fmt.Printf("백두산의 주소는 %s입니다.\n", m["백두산"])
}
```

출력문

```go
송하나의 주소는 서울시 강남구입니다.
백두산의 주소는 부산시 사하구입니다.
```

- make() 함수로 맵 생성
- make[key]value

```go
package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product) 

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}

	for k, v := range m { 
		fmt.Println(k, v)
	}
}
```

출력문

```go
46 {지우개 200}
78 {자 1000}
345 {샤프 3000}
897 {샤프심 500}
16 {볼펜 500}
```

- range 키워드로 순회 : 첫 번째 값 키, 두 번째 값 값
- 맵 내부 요소 보관 : 입력 순서, 키 값 상관없음

### 22.3.1 요소 삭제와 없는 요소 확인

```go
delete(m, key)  //맵 변수, 삭제 키
```

```go
package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)
	fmt.Println(m[3])
	fmt.Println(m[1])
}
```

출력문

```go
0
0
```

- 없는 요소 삭제 시도 → 아무 동작 X
- 이미 삭제된 요소 출략 → 타입 기본값 (int : 0)
- 값이 0일 때와 요소 없을 때 → 복수 반환 (값, 요소 존재 여부)

```go
v, ok := m[3]
```

### 22.3.2 맵, 배열, 리스트 속도 비교

| 행위 | 배열, 슬라이스 | 리스트 | 맵 |
| --- | --- | --- | --- |
| 추가 | O(N) | O(1) | O(1) |
| 삭제 | O(N) | O(1) | O(1) |
| 읽기 | O(1) - 인덱스로 접근 | O(N) - 인덱스로 접근 | O(1) - 키로 접근 |

## 연습문제

1. 

| 행위 | 배열, 슬라이스 | 리스트 | 맵 |
| --- | --- | --- | --- |
| 추가 | O(N) | O(1) | O(1) |
| 삭제 | O(N) | O(1) | O(1) |
| 읽기 | O(1) - 인덱스로 접근 | O(N) - 인덱스로 접근 | O(1) - 키로 접근 |

2.

(1) map[int]Student

(2) 23

(3) 38

3.

(1) 큐

(2) 맵

(3) 이중 배열

(4) 링