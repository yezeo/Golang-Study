# Chapter 13. Structure

## 13.1 선언 및 기본 사용

구조체란? 서로 다른 타입의 값들을 변수 하나로 묶어주는 기능    

### 정의 방법

```go
type 타입명 struct {
	필드명 타입
	...
	필드명 타입
}
```

예시

```go
type Student struct {
	Name string
	Class int
	StudentNo int
	Score float64
}
```

→ 이제 Student 타입을 int나 float64 같은 내장 타입처럼 선언해 사용할 수 있다.   

Student 타입의 구조체 변수 a를 선언하려면 → `var a Student` 

### 구조체 정의, 사용 예제

```go
package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var myHouse House
	myHouse.Address = "Gwangmyeong"
	myHouse.Size = 32
	myHouse.Price = 9.8
	myHouse.Type = "Apartment"

	fmt.Println("주소: ", myHouse.Address)
	fmt.Printf("크기: %d평", myHouse.Size)
	fmt.Printf("가격: %.2f억 원\n", myHouse.Price)
	fmt.Println("타입: ", myHouse.Type)
}
```

## 13.2 구조체 변수 초기화

방법은 3가지가 있다.    

- 초깃값 생략
- 모든 필드 초기화
- 일부 필드만 초기화

### 📌 초깃값 생략

초깃값 생략 시, 모든 필드가 기본값으로 초기화된다.

```go
var house House
```

→ 각 필드의 타입에 따라 기본값이 부여된다.    

- house.Address 는 " " (빈 문자열)
- house.Size 는 0
- house.Price 는 0.0

### 📌 모든 필드 초기화

모든 필드 값을 중괄호 { } 사이에 넣어서 초기화한다. (배열과 비슷함)

```go
var house1 House = House{"서울시", 28, 9.8, "아파트"}

var house2 House = House{
	"광명시",
	32,
	14.0,
	"주택",  // 여러 줄로 초기화 할 때는 제일 마지막 값 뒤에 꼭 쉼표를 달아주기!
}
```

### 📌 일부 필드만 초기화

'필드명:필드값' 형식으로 초기화한다.    

초기화하지 않은 나머지 변수에는 기본값이 할당된다.

```go
var house House = House{Size:28, Type:"아파트"}
```

## 13.3 구조체를 포함하는 구조체

구조체의 필드로 다른 구조체를 포함할 수도 있다. 방식은 아래 두 가지.    

- 내장 타입처럼 포함하는 방식
- 포함된 필드 방식

### 📌 방식 1: 내장 타입처럼 포함하는 방식

```go
type User struct {
	Name string
	ID string
	Age int
}
type VIP struct {
	UserInfo User   // User 구조체를 포함
	Level int
	Price int
}
```

사용은 이런 식으로 하도록 하자.

```go
user := User{"Yoonseo", "dotsi", 22}
vip := VIP{
		User{"Eunxung", "ae", 22},
		3,
		250,
}
fmt.Println(vip.UserInfo.Name) 
		// Eunxung 출력
```

### 📌 방식 2: 포함된 필드 방식

위와 다른 점은, vip 선언할 때 아래와 같이 User의 필드명을 생략했다는 점이다. 

```go
type VIP struct {
	User   // 필드명 생략
	Level int
	Price int
}
```

이 때 더 쉽게 사용할 수 있다.    

위에서 `vip.UserInfo.Name`으로 불러왔던 것을, `vip.Name`으로 점 하나로 바로 접근할 수 있다.   

이를 '포함된 필드'라고 부른다. 포함된 필드를 이용하면 점 . 을 두 번 찍을 필요 없이 한 번만으로 바로 접근할 수 있어서 더 편리하다.

### → 이 때, 필드명이 겹치면? 어떻게 해결하는지 알아보자.

포함된 필드 안에도 Level 필드가 있고, 상위 구조체에도 Level 필드가 있다면 어떻게 할까?

```go
package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int // User의 Level 필드
}

type VIPUser struct {
	User  // Level 필드를 갖는 구조체
	Price int
	Level int // VIPUser의 Level 필드
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 10},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d 유저 레벨:%d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,      // VIPUser의 Level
		vip.User.Level, // 포함된 구조체명을 쓰고 접근
	)
}
```

위와 같이,    

상위 구조체의 필드는 `vip.Level`,     

포함된 필드 구조체의 필드는 `vip.User.Level` 처럼 포함된 구조체명을 쓰고 접근하면 된다.   

## 13.4 구조체 크기

구조체가 차지하는 메모리 크기는 어떻게 알 수 있을까? 구조체 크기를 구해보자.

```go
type User struct {
	Age int
	Score float64
}
```

위와 같은 구조체 User가 정의되어 있을 때,    

```go
var User user
```

User 구조체의 변수 user가 선언되면, 컴퓨터는 Age와 Score 필드를 **연속되게** 담을 수 있는 메모리 공간을 찾아 할당한다.   

int 타입은 8바이트, float64도 8바이트이므로, 총 **16바이트** 크기가 필요하다.    

즉, 구조체 변수 user의 크기는 16바이트가 된다. → 따라서 User도 16바이트.

### 📌 구조체 값 복사

구조체 변숫값을 다른 구조체에 대입하면 모든 필드값이 복사된다.

```go
type Student struct {
	Age   int 
	No    int
	Score float64
}

func main() {
	var student = Student{15, 23, 88.2}

	student2 := student   // student 구조체 모든 필드가 student2 로 복사된다.
```

# Chapter 13 연습문제

1.  

```go
type Product struct {
	Name string
	Price int
	ReviewScore float64
}
```

2.  

```go
200
8.7
```

3.  패딩을 최대한 줄인 구조체 크기= 32바이트

```go
type Padding struct {
	A int8   //1byte
	G int8   //1byte
	D uint16  //2byte
	F float32 //4byte
	B int    //8byte
	C float64 //8byte
	E int    //8byte
}
// 1+1+2+4+8+8+8 = 32
```
