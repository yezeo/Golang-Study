# Chapter 20 인터페이스

## 20.1 인터페이스

인터페이스 : 상호작용면? → 메서드 구현을 포함한 구체화된 객체가 아닌 추상화된 객체로 상호작용

### 20.1.1 인터페이스 선언

type + 인터페이스명 + interface 키워드 + { 메서드 집합 }

```go
type DuckInterface interface {
	Fly()
	Walk(distance int) int
}
```

type 키워드? 인터페이스 변수 선언 가능, 변수의 값으로 사용 가능

메서드 집합 유의사항

1. 메서드는 반드시 메서드명이 있어야 함
2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없음
3. 인터페이스에서는 메서드 구현을 포함하지 않음

```go
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {

	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())
}
```

출력문

```go
안녕! 나는 12살 철수라고 해
```

- Stringer 인터페이스 선언, 매새변수 없이 string 타입 반환하는 String() 메서드 포함
- Student 타입 String() 메서드 포함 → Student 타입은 Stringer 인터페이스로 사용 가능
- fmt 패키지의 Sprintf() 함수: 서식에 따라 문자열을 만들어서 반환
- stringer값으로 Student 타입 변수 student를 대입
- stringer 인터페이스가 가지고 있는 메서드 String() 호출 → stringer값으로 Student 타입 student를 가지고 있기 때문에 student의 메서드 String() 호출

## 20.2 인터페이스 왜 쓰나?

인터페이스를 이용하면 구체화된 객체가 아닌 인터페이스만 가지고 메서드를 호출할 수 있기 때문에 큰 코드 수정 없이 필요에 따라 구체화된 객체를 바꿔서 사용할 수 있게 됩니다. 그럼으로써 프로그램의 변경 요청에 유연하게 대처할 수 있게 됐습니다.

온라인 쇼핑몰 택배 전송 프로그램 : Fedex와 우체국에서 제공하는 각 패키지 → 인터페이스

```go
// Fedex에서 제공한 패키지
package fedex

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
```

```go
package main

import "github.com/tuckersGo/musthaveGo/ch20/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
```

출력문

```go
Fedex sends 어린 왕자 parcel
Fedex sends 그리스인 조르바 parcel
```

```go
// 우체국에서 제공한 패키지
package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다.\n", parcel)
}
```

```go
package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
```

출력문

```go
에러 발생
```

- sender 변수는 *koreaPost.PostSender 타입
- SendBook() 함수의 인수 : *fedex.FedexSender → sender 타입이 달라 에러 발생

→ 인터페이스 활용해서 문제 해결

```go
package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}
```

출력문

```go
우체국에서 택배 어린 왕자를 보냅니다.
우체국에서 택배 그리스인 조르바를 보냅니다.
Fedex sends 어린 왕자 parcel
Fedex sends 그리스인 조르바 parcel
```

- Sender 인터페이스 : Send() 메서드 포함, SendBook() 함수 인수 Sender 인터페이스 → SendBook() 함수 내부에서 Send() 메서드 호출
- PostSender, FedexSender 모두 Send(string) 메서드 포함 → Sender 인터페이스 사용 가능 = SendBook() 함수 인수로 사용 가능
- SendBook() 함수는 어떤 타입이든지 상관없이 Send() 인덱스만 제공하면 OK

### 20.2.1 추상화 계층

추상화 : 내부 동작을 감춰서 서비스 제공자와 사용자 모두에게 자유를 주는 방식

인터페이스는 추상화를 제공하는 **추상화 계층**

예) Sender 인터페이스 : 택배 서비스 사용자와 제공자 서로 구현에 신경 쓰지 않고 추상화된 관계 중심으로 코딩

디커플링 : 추상화 계층을 이용해 서로 결합을 끊는 것 (결합도는 낮출수록 좋다)

## 20.3 덕 타이핑

덕 타이핑 : 타입 선언 시 구현 여부를 명시적으로 나타낼 필요 없이 인터페이스에 정의한 메서드 포함 여부만으로 결정하는 방식

### 20.3.1 서비스 사용자 중심 코딩

덕타이핑의 장점 : 서비스 사용자 중심의 코딩 가능

인터페이스가 사용될 때 해당 타입이 인터페이스에 정의된 메서드를 포함했는지 여부로 결정

## 20.4 인터페이스 기능 더 알기

기본 기능 외 3가지 다른 기능

- 포함된 인터페이스
- 빈 인터페이스
- 인터페이스 기본값

### 20.4.1 인터페이스를 포함하는 인터페이스

구조체처럼 인터페이스도 다른 인터페이스를 포함할 수 있다

```go
type Reader interface {
	Read() (n int, err error)
	Close() error	
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}
```

- ReadWriter 인터페이스는 Reader와 Writer 인터페이스 메서드 집합 포함
- 하나의 Close() error 메서드만 포함
1. Read(), Write(), Close() 메서드를 포함한 타입 → Reader, Writer, ReadWriter 모두 사용 가능
2. Read(), Close() 메서드를 포함한 타입 → Reader만 사용 가능
3. Write(), Close() 메서드를 포함한 타입 → Writer만 사용 가능
4. Read(), Write() 메서드를 포함한 타입 → Close() 메서드가 없기 때문에 모두 사용 불가능

### 20.4.2 빈 인터페이스 interface{}를 인수로 받기

interface{} : 메서드를 가지고 있지 않은 빈 인터페이스

모든 타입이 빈 인터페이스로 쓰일 수 있다. 어떤 값이든 받을 수 있는 함수, 메서드, 변숫값을 만들 때 사용한다

```go
package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)          // int
	PrintVal(3.14)        // float64
	PrintVal("Hello")     // string
	PrintVal(Student{15}) // Student
}
```

출력문

```go
v is int 10
v is float64 3.140000
v is string Hello
Not supported type: main.Student:{15}
```

### 20.4.3 인터페이스 기본값 nil

인터페이스 변수의 기본값 : nil (유효하지 않은 메모리 주소)

```go
package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker 
	att.Attack()     
}
```

출력문

```go
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0x59ab56]     

goroutine 1 [running]:
main.main()
        C:/Users/WINDOWS10/Desktop/Golang-Study/Ch20/ex20.6/ex20.6.go:9 +0x16
```

- 변수 att는 Attacker 인터페이스, 초깃값이 없으므로 nil
- att.Attack()에서 att값이 nil이기 때문에 실행 중 에러 발생 (런타임 에러)

## 20.5 인터페이스 변환하기

- 구체화된 다른 타입으로 타입 변환하기
- 다른 인터페이스로 타입 변환하기

### 20.5.1 구체화된 다른 타입으로 타입 변환하기

인터페이스를 본래의 구체화된 타입으로 복원할 때 주로 사용

```go
var a Interface
t := a.(ConcreteType)
```

- a를 ConcreteType 타입으로 변경 → t에 반환

```go
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {

	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)

}

func main() {
	s := &Student{15}

	PrintAge(s)
}
```

출력문

```go
Age: 15
```

- Stringer 인터페이스는 String() 메서드 포함
- 구조체 포인터 *Student 타입은 String() 메서드 포함 → PrintAge() 함수의 stringer 인터페이스 인수로 사용 가능
- *Student 타입 변수 s, Age = 15
- PrintAge() 함수 : Stringer 인터페이스 변수 → *Student 타입 (stringer 인스턴스 변수 내부에 *Student 타입 인스턴스를 가리키고 있음 → Age 접근 가능

### 20.5.2 다른 인터페이스로 타입 변환하기

인터페이스가 변경 전 인터페이스를 포함하지 않아도 된다. 하지만 인터페이스가 가리키고 있는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야 한다

```go
var a AInterface = ConcreteType{}
b := a.(Binterface)
```

- ConcreteType이 AInterface와 BInterface를 모두 포함

```go
package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	c := reader.(Closer)
	c.Close()
}

func main() {
	file := &File{}
	ReadFile(file)
}
```

출력문

```go
panic: interface conversion: *main.File is not main.Closer: missing method Close

goroutine 1 [running]:
main.ReadFile(...)
        C:/Users/WINDOWS10/Desktop/Golang-Study/Ch20/ex20.10/ex20.10.go:18
main.main()
        C:/Users/WINDOWS10/Desktop/Golang-Study/Ch20/ex20.10/ex20.10.go:24 +0x27
```

- *File은 Read() 메서드 포함 → Reader 인터페이스로 사용 가능
- Reader 인터페이스 변수 → Closer 인터페이스 : 서로 다른 메서드 집합을 가지고 있어도 문법적 문제 X
- 문제는 reader 인터페이스 변수가 *File 타입을 가리키고 *File 타입은 Close() 메서드 포함 X = Closer 인터페이스로 사용 불가능

### 20.5.3 타입 변환 성공 여부 반환

타입 변환 반환값을 두 개의 변수로 받기, 두 번째 반환값 = 타입 변환 가능 여부

```go
var a Interface
t, ok := a.(ConcreteType)
```

## 연습문제

1. *File 타입은 Read() 메서드만 가지고 있어서 Read()와 Write()가 정의된 ReadWriter 인터페이스로 사용할 수 없다.
2. 

```go
type DB interface {
	GetData() string
	WriteData(data string)
	Close() error
}
```

3.

```go
func CheckAndRun(stringer Stringer) {
	if r, ok := stringer.(Reader); ok{
		r.Read()
	}
}
```