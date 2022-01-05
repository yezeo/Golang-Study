# Chapter 22. 자료구조

## 22.1 리스트

배열 VS 리스트

- 배열: 연속된 메모리에 데이터 저장
- 리스트: 불연속된 메모리에 데이터 저장 (포인터로 연결)

### 📌 리스트는 요소들을 `포인터`로 연결한 자료구조이다.

```go
type Element struct {  // 구조체
	Value interface{ }   // 데이터
	Next *Element        // 다음 요소의 주소를 저장하는 포인터
	Prev *Element        // 이전 요소의 주소를 저장하는 포인터
}
```

### 📌 리스트 기본 사용법

- 리스트 **생성**: `list.New()`
- 리스트 **요소 추가**
    - `v.PushBack(e)` , `v.PushFront(e)`
    - `v.InsertBefore(e, pos)` , `v.InsertAfter(e, pos)`
- 리스트 **순회**: `for e:=v.Front(); e!=nil; e=e.Next()`

```go
package main
import (
	"container/list"
	"fmt"
)
func main(){
	v := list.New()  //리스트 생성
	e4 := v.PushBack(4)  //리스트 뒤에 요소 추가
	e1 := v.PushFront(1)  // 리스트 앞에 요소 추가 -> {1 4}
	v.InsertBefore(3, e4) // e4 요소 앞에 3 추가 -> {1 3 4}
	v.InsertAfter(2, e1)  // e1 요소 뒤에 2 추가 -> {1 2 3 4}

	for e := v.Front(); e != nil; e=e.Next(){
		fmt.Print(e.Value, " ")
	}
fmt.Println()
for e:= v.Back(); e != nil; e=e.Prev()
}
```

[📌 배열 VS 인덱스 (Big-O 시간복잡도 비교)](https://www.notion.so/b9d2347a617d4cd4ac1d441d7056245a)

## 리스트 응용 1) 큐(Queue) 구현하기

```go
package main
import (
	"fmt"
	"container/list"
)
type Queue struct {  // queue 구조체 정의
	v *list.List
}
func (q *Queue) Push(val interface{}){
	q.v.PushBack(val)
}
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil //요소 없을 경우
}
func NewQueue() *Queue {
	return &Queue{ list.New() }
}
func main(){
	queue := NewQueue()

	for i:=1; i<5; i++ {
		queue.Push(i)  // {1,2,3,4}
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v) // 1->2->3->4->
		v = queue.Pop()
	}
}
```

## 리스트 응용 2) 스택(Stack) 구현하기

```go
package main

import (
	"fmt"
	"container/list"
)

type Stack struct {
	v *list.List
}
func NewStack() *Stack {
	return &Stack{ list.New() } 
}
func (s *Stack) Push(val interface{}){
	s.v.PushBack(val)
}
func (s *Stack) Pp() interface{} {
	back := s.v.Back() //Queue와 다른 부분, Back부터 remove
	if back != nil {
		return s.v.Remove(back)
	
	return nil 
}
func main(){   
	stack := NewStck()
	fr i:=1; i<5; i++{
		stack.Push(i)
	}
	val := stack.Pop()
	for val != nil {
		mt.Printf("%v -> ", val)  // 4->3->2->1->
		val = stack.Pop()

	}
}
```

## 22.2 링

링(ring)이란? 맨 뒤-맨 앞 요소가 서로 연결된 원형 자료구조.

- 시작도 없고, 끝도 없음. 현재 위치만 알고 있음.
- ‘환형 리스트’라고도 불림.

```go
package main
import (
	"container/ring"
	"fmt"
)
func main(){
	r:=ring.New(5) //요소가 5개인 링 생성

	n := r.Len() //링 길이 반환 

	for i:=0; i<n; i++ {
		r.Value = 'A' + i  // A B C D E
		r = r.Next()
	}
	for j:=0; j<n; j++ {
		fmt.Printf("%c ", r.Value)  //순회하며 값 출력
		r = r.Next()
	}
	fmt.Println()

	for j:=0; j<n; j++ {  
		fmt.Printf("%c ", r.Value)  // 역순회하며 값 출력
		r = r.Prev()
	}
}
```

### 📌 링은 언제 쓸까?

링은 저장할 개수가 고정되고, 오래된 요소는 지워도 되는 경우에 적합하다. 

예를 들어 아래와 같은 경우에 사용된다. 

1. **실행 취소** 기능
2. **고정 크기 버퍼** 기능
3. **리플레이** 기능

## 22.3 맵

맵(map)이란? 키와 값 형태로 데이터를 저장하는 자료구조.   

리스트나 링과 달리, container 패키지를 임포트할 필요가 없는 Go 기본 내장 타입.

### 📌 map 기본 사용법

- 맵 생성: `make( map[key]value )`

```go
package main
import "fmt"

func main(){
	m := make(map[string]string)
	m["김윤서"] = "경기도 광명시"
	m["이선주"] = "대전광역시 봉명동"
	m["윤창용"] = "서울특별시 염창동"

	m["윤창용"] = "서울특별시 상암동"

	fmt.Printf("김윤서의 주소는 %s입니다.", m["김윤서"])
	fmt.Printf("윤창용의 주소는 %s입니다.", m["윤창용"])
}
```

- 맵 순회: `for k,v := range m { }`

```go
package main
import "fmt"

type Product struct {
	Name string
	Price int
}
func main(){
	m := make(map[int]Product)  //key가 int, value가 구조체인 맵
	
	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"필통", 1000}
	m[345] = Product{"샤프", 700}
	
	for k,v := range m {
		fmt.Println(k,v)
	}
}
```

- 요소 삭제: `delete(m, key)`
- 없는 요소 확인: `v, ok := m[3]`  //ok에 존재 여부가 boolean으로 저장됨

[📌 맵, 배열, 리스트 속도 비교](https://www.notion.so/cf08c80bf0884c5cb3ca6b61a1232e36)
