# Chapter 15. String

## 15.1 문자열

Go의 문자열 출력 방법은 2가지가 있다.   

- 큰따옴표
- 백쿼트

```go
package main

import "fmt"

func main() {
	
// 1) 큰따옴표로 묶기
	str1 := "Hello\t'World'\n"
	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를\n"
	
// 2) 백쿼트로 묶기
	str2 := `Hello\t"Go"\nWorld!`
	poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println()
	fmt.Println(poet1)
	fmt.Println(poet2)
}
```

→ 결과   

백쿼트로 묶을 경우, 특수문자인 `\t` `\n`이 그대로 출력된다.   

또한 백쿼트는 여러 줄에 걸쳐서 문자열을 쓸 수 있지만, 큰따옴표는 한 줄에만 쓸 수 있어 개행하려면 `\n`을 사용해야 한다.

```bash
C:\Users\a\Desktop\goproject\chapter 15\ex15.1>.\ex15.1.exe

Hello   'World'

Hello\t"Go"\nWorld!

죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를

죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를
잎새에 이는 바람에도
나는 괴로워했다.
```

### 📌 문자 한 개 담기; `rune` 타입

문자 하나를 표현하는 데 rune 타입을 사용한다.    

**rune** 타입은 **int32** 타입의 별칭 타입이다. 즉, 이름만 다를 뿐 int32와 성질이 같다.    

문자를 표현하기 위해 **별칭 rune**을 사용한다고 이해하자.

```go
package main

import "fmt"

func main() {
	var ch1 rune = '음'

	fmt.Printf("%T\n", ch1) // ch1의 타입 출력
	fmt.Println(ch1)        // ch1의 값 출력
	fmt.Printf("%c\n", ch1) // ch1에 저장된 값을 문자로 출력
}
```

→ 결과    

rune으로 정의한 문자 ch1의    

- **타입** `%T` 를 출력하면 → `int32`
- **값**을 출력하면 → `51020` (int32 타입이므로 값은 숫자가 나온다)
- **문자** `%c` 을 출력하면 → `음` (문자 그대로)

```go
int32
51020
음
```

### 📌 문자열 크기 알아내기; `len()`

len()을 통해 출력되는 '문자열 크기'는 말 그대로 문자열이 차지하는 메모리 크기이다.   

한글은 한 글자 당 3 byte, 영어는 한 글자 당 1 byte를 차지한다.   

```go
package main

import "fmt"

func main() {
	str1 := "가나다라마"
	str2 := "abcde"

	fmt.Printf("'가나다라마'의 크기= %d\n", len(str1)) //한글 문자열 크기
	fmt.Printf("'abcde'의 크기= %d\n", len(str2)) //영문 문자열 크기
	
}
```

→ 결과

```bash
'가나다라마'의 크기= 15
'abcde'의 크기= 5
```

### 📌 글자 수 알아내기; `[]rune` 타입 변환

`[]rune` 타입은 **상호 타입 변환**이 가능하다. (**슬라이스 타입**이라서, 이는 18장에서 설명)   

rune 배열의 각 요소에는 문자열의 각 글자가 대입된다.

[]rune 타입을 이용해 문자열의 글자 수를 알아내는 방법은   

1. 타입 변환 (string을 []rune 타입으로)  `runes := []rune(str)`
2. `len(runes)` ⇒ 글자 수가 반환된다.   

```go
package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str) // 문자열을 []rune 타입으로 변환

	fmt.Printf("len(str) = %d\n", len(str))     // 메모리 크기 반환
	fmt.Printf("len(runes) = %d\n", len(runes)) //글자수 반환; len(runes)
}
```

→ 결과

✔ `len(**[]rune(str)**)` 이렇게도 가능

```bash
len(str) = 12
len(runes) = 8
```

### 💡 또한, String 타입을 []byte 로 타입 변환할 수 있다.

→ 20장 '인터페이스' 와 A.3절 '입출력 처리'에서 다룬다.    

지금은 string과 []byte 타입 간 상호 변환이 가능하다는 것만 알고 넘어가자.

## 15.2 문자열 순회

1. **인덱스**를 사용한 바이트 단위 순회
2. `**[]rune**` 으로 변환 후 한 글자씩 순회
3. **`range`** 를 이용해 한 글자씩 순회

### 1. 인덱스를 사용한 바이트 단위 순회하기

```go
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T\t값:%d\t문자값:%c\n", str[i], str[i], str[i])
	}
```

→ 결과

```go
타입: uint8     값:72   문자값:H
타입: uint8     값:101  문자값:e
타입: uint8     값:108  문자값:l
타입: uint8     값:108  문자값:l
타입: uint8     값:111  문자값:o
타입: uint8     값:32   문자값:
타입: uint8     값:236  문자값:ì
타입: uint8     값:155  문자값:
타입: uint8     값:148  문자값:
타입: uint8     값:235  문자값:ë
타입: uint8     값:147  문자값:
타입: uint8     값:156  문자값:
타입: uint8     값:33   문자값:!
```

- 인덱스를 사용해 각 바이트값을 출력한다.
- 한글이 깨진다. 왜? 영어는 1byte이므로 잘 출력되지만, 한글은 3byte이므로 깨진다.

### 2. `[]rune` 으로 타입 변환 후 한 글자씩 순회하기

```go
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T\t값:%d\t문자값:%c\n", arr[i], arr[i], arr[i])
	}
```

→ 결과

```go
타입: int32     값:72     문자값:H
타입: int32     값:101    문자값:e
타입: int32     값:108    문자값:l
타입: int32     값:108    문자값:l
타입: int32     값:111    문자값:o
타입: int32     값:32     문자값:
타입: int32     값:50900  문자값:월
타입: int32     값:46300  문자값:드
타입: int32     값:33     문자값:!
```

- 한글도 잘 출력된다.
- 그러나, []rune 으로 변환하는 과정에서 불필요한 메모리 할당(arr 정의)이 일어난다.
- 이를 해결하려면, range 를 사용하자.

### 3. `range` 키워드 사용하여 한 글자씩 순회

```go
	str := "Hello 월드!"
	for _, v := range str {
		fmt.Printf("타입: %T\t값:%d\t문자값:%c\n", v, v, v)
	}
```

→ 결과

```go
타입: int32     값:72     문자값:H
타입: int32     값:101    문자값:e
타입: int32     값:108    문자값:l
타입: int32     값:108    문자값:l
타입: int32     값:111    문자값:o
타입: int32     값:32     문자값:
타입: int32     값:50900  문자값:월
타입: int32     값:46300  문자값:드
타입: int32     값:33     문자값:!
```

- 이처럼 range를 이용하면, 추가 메모리 할당 없이 문자열을 한 글자씩 순회할 수 있다.
- 모든 문자 타입이 int32, 즉 rune 이다.
- 한글, 영어 모두 잘 출력 된다.

## 15.3 문자열 합치기

`+`와 `+=` 연산을 사용해서 문자열을 이을 수 있다.   

```go
package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2 // str1 + " " + str2 와 동일
	fmt.Println(str1)

}
```

### 📌 문자열 비교하기

연산자 `==` , `!=` 을 사용해서 두 문자열이 같은지 비교할 수 있다.

```go
package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "Hell"
	str3 := "Hello"

	fmt.Printf("%s == %s  : %v\n", str1, str2, str1 == str2)
	fmt.Printf("%s != %s : %v\n", str1, str2, str1 != str2)
	fmt.Printf("%s == %s : %v\n", str1, str3, str1 == str3)
	fmt.Printf("%s != %s : %v\n", str1, str3, str1 != str3)
}
```

### 📌 문자열 대소 비교하기

`>` , `<` , `>=`, `<=` 연산자를 이용해서 문자열 간 대소 비교를 할 수 있다.    

문자열 대소 비교는 첫 글자부터 하나씩 값을 비교해서, 그 글자의 유니코드 값이 다를 경우 대소를 반환한다.    

즉, 첫 글자가 더 큰 값이면, 더 큰 문자열이다.

```go
package main

import "fmt"

func main() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)   
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)   
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4) 
}
```

→ 결과

```go
BBB > aaaaAAA : false
BBB < BBAD : false
BBB <= ZZZ : true
```
