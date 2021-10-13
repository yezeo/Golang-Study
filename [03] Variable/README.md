
## 3.1 변수 선언

- **변수**란???? : 값을 저장하는 메모리 공간
- 변수 선언 방법
    - `var a int = 10`
    - `var b int`  // 초깃값 생략 (기본값으로 대체)
    - `var c = 4`  // 타입 생략 (우변 값의 타입이 됨)
    - `d := 10`  // 선언 대입문 `:=`을 사용해서 var 키워드와 타입 생략
- 변수의 원리

`var minimumWage int = 10` 처럼 변수를 선언하면, 

1) 메모리에 저장할 공간을 만들고, (int만큼)

2) minimumWage 라고 지칭한 뒤,

3) 값 10을 복사한다.

→ 이제 minimumWage 라는 변수명을 이용해서 **해당 공간에 접근할 수 있다**.

## 3.2 변수 특징

### 📌 타입별 기본값

- 모든 **정수** 타입(*int, uint, byte, rune, int8, uint64등등*)은  **0**
- 모든 **실수** 타입(*float32, float64, complex64등등*)은 **0.0**
- 불리언은 **false** 가 기본값,
- 문자열은 **""** (빈 문자열)
- 그 외에는 **nil** (정의되지 않은 메모리주소를 나타내는 Go 키워드) 로 기본값이 설정된다.
    
    (거의 다른 언어의 기본값과 비슷하다!!)
    

### 📌 기본 타입

타입 생략 시, 우변의 타입으로 변수 타입이 자동 지정됨.

- 정수 값의 기본 타입은 **int**
- 실수 값의 기본 타입은 **float64**

```go
var b = 3  //b의 타입은 int
var c = 0.001  //c의 타입은 float64
```

## 3.3 타입 변환과 주의점

Go는 **최강 타입 언어**이다! → 융통성이 매우 없다는 뜻.

쉽게 말해, 연산이나 대입에서 타입이 다르면 에러가 발생한다.

심지어 int와 int64 값 사이 호환도 안 된다 → 개발자가 하나하나 타입 변환을 해줘야 함

```go
var a int
var b int64 = 3
a = b  // ERROR. a는 int타입, b는 int64타입이므로 대입연산 불가.

var c float64 = 3.5
d := a*c  // 당연히 ERROR. a는 int, c는 float형이므로 연산 불가.
```

다른 언어에서는 자동으로 변환해주는 타입들도 Go는 지원하지 않는다.

### → 개발자가 직접 타입 변환 해줘야 한다.

```go
a = int(b)
d := float64(a) * c
```

원하는 타입명을 적고 ( )로 변수를 묶어준다. 

---

### 📌 주의점 1: 실수→정수 타입 변환

**실수 → 정수** 타입으로 변환하면, 소수점 이하 숫자가 없어진다.

```go
a := 3.5
var b int = int(a * 3) // b= int(10.5) = 10
var c int = int(a) * 3 // c= 3*3 = 9
var isEqual bool = (b == c) // 두 값이 같은가

fmt.Println(isEqual) // false
fmt.Println(b, c)  // 10 9
```

→ 결과   
![Untitled (1)](https://user-images.githubusercontent.com/61778930/136971477-74680d6c-d82e-4498-8b0b-6a5f9dee8067.png)


### 📌 주의점 2: **큰 범위→작은 범위 타입** 변환

**큰 범위 → 작은 범위** 타입 변환 시, 바이트 변화로 값이 전혀 달라질 수 있다.

ex. int16 → int8 로 변환할 때 상위 1바이트가 없어져 값이 변화

```go
var typeBig int16 = 3456
var typeSmall int8 = int8(typeBig)

fmt.Println(typeBig)
fmt.Println(typeSmall)
```

→ 결과

![Untitled (2)](https://user-images.githubusercontent.com/61778930/136971506-bd60d93f-9229-4621-8e37-c8c283cd666e.png)


---

### 변수의 범위: 전역 변수, 지역 변수

변수는 자신이 속한 중괄호 {} 범위를 벗어나면 사라진다.

```go
package main

import "fmt"

var g int = 10 //global variable

func main() {
	var m int = 20 //local variable
	{
		var s int = 50 //local variable in {}
		fmt.Println(m, s, g)
	
	} // variable 's' dead

	fmt.Println(m, g)  // 20 10
	fmt.Println(s)  // ERROR
}
```

끝.

---

### 다시 짚어보는 실행 명령어~~~

1. `cd ottl`
2. `go mod init goproject/ottl`
3. `go build`
4. `.\ottl.exe`
