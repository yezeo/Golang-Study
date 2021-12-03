# Chapter 18. Slice

## 18.1 슬라이스

슬라이스란? → GO 언어에서 제공하는 **동적 배열**.    

즉, 자동으로 배열 크기를 증가시키는 자료구조이다.   

### 📌 슬라이스 선언 방법

```go
var slice []int
```

일반적인 배열은 `var array [10]int` 와 같이 [ ] 안에 크기를 선언해줘야 하는 반면, 슬라이스는 적지 않고 선언할 수 있다.

그러나, 슬라이스를 초기화하지 않고 사용하면, 길이가 0인 슬라이스가 만들어진다.    

슬라이스 길이를 초과하게 되어 접근하게 되면 런 타임 에러가 발생한다. 아래 예시를 보자.

```go
package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 { // slice 길이가 0인지 확인
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10   // ERROR (패닉 발생)
	fmt.Println(slice)
}
```

그렇다면, 슬라이스를 초기화하는 방법에 대해 알아보자.

### 초기화 방법 1: `{ }` 이용

```go
var slice1 = []int{1,2,3}
var slice2 = []int{1, 5:2, 10:3}  // 1 0 0 0 0 2 0 0 0 0 3
```

### 💡 참고)

아래 두 구문은 서로 다른 타입을 만든다. 첫 번째는 길이가 3인 고정 길이 배열을 만들고 두 번째는 슬라이스를 만든다.    

배열과 슬라이스는 엄연히 다른 자료구조이다.

```go
var array = [...]int{1,2,3}
var slice = []int{1,2,3}    
```

### 초기화 방법 2: `make()` 이용

make() 함수의 첫 번째 인수로 슬라이스의 **타입**을 적어주고, 두 번째 인수로 **길이**를 적어준다.   

```go
var slice = make([]int, 3)
```

→ 결과, `{0, 0, 0}`의 값을 갖는 슬라이스가 생성된다 (int 타입 기본값인 0으로 초기화되었다)

### 👀 초기화한 후에는 슬라이스 요소에 접근할 수 있다.

```go
var slice = make([]int, 3)
slice[1] = 5
```

### 📌 슬라이스 순회

슬라이스 순회 역시 배열과 같다. 동적으로 길이가 늘어나는 점만 제외하면 대부분 배열과 비슷하다.

```go
var slice = []int{1,2,3}

for i := 0; i<len(slice); i++ {  // 각 요소에 10 더하기
	slice[i] += 10
}
for i, v := range slice {  //각 요소에 2 곱하기
	slice[i] = v*2
}
```

### 📌 슬라이스 요소 추가: `append()`

기존 배열은 한 번 길이가 정해지면 늘릴 수 없지만, 슬라이스는 요소를 추가함으로써 길이를 늘릴 수 있다.    

슬라이스 요소 추가에 append() 함수를 사용한다.    

```go
var slice = []int{1,2,3}
slice2 := append(slice, 4) // 요소 추가, {1,2,3,4}가 된다.
```

첫 번째 인수로 슬라이스를 넣어주고, 그 뒤로 추가하고자 하는 요소를 넣어준다.    

실행 시, 슬라이스의 맨 뒤에 요소를 추가할 수 있다.    

- 여러 값을 추가할 수도 있다.

```go
slice = append(slice, 3,4,5,6,7)
```

## 18.3 슬라이싱

**슬라이싱**은 배열의 일부를 집어내는 기능을 말한다.

```go
array[시작인덱스:끝인덱스]
```

예제를 통해 살펴보자.

```go
package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2] // 슬라이싱

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100 // array의 두 번째 값 변경

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500) // slice에 값 추가

	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}
```

## 18.5 슬라이스 정렬

Go 언어에서 기본 제공하는 sort 패키지를 사용해 정렬할 수 있다.

### 📌 int 슬라이스 정렬

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // 정렬되지 않은 슬라이스
	sort.Ints(s)                 // 정렬
	fmt.Println(s)
}
```

### 📌 구조체 슬라이스 정렬

```go
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// []Student의 별칭 타입 Students
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age } // 나이 비교
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18}}

	sort.Sort(Students(s)) // 정렬
	fmt.Println(s)
}
```
