# Chapter 23. 에러 핸들링

- 프로그램의 에러를 처리하는 방법

## 23-I. 에러 반환

```go
if err != nil {
    return err
}
```

### 23-I-1. 사용자 에러 반환

인수로 문자열을 입력하면 인수와 같은 메시지를 갖는 error 생성하여 반환함

```go
func New(text string) error
```

다음과 같이 사용 가능

```go
import "errors"

errors.New("에러 미시지")
```

### 23-I-2. 에러 랩핑

에러를 감싸서 새로운 에러를 만들어야 할 경우 존재

```go
if err != nil {
    return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s error: %w", word, err)
}
```

감싸진 에러를 다시 꺼내올 때는 errors 패키지의 As() 함수 이용

```go
if errors.As(err, &numError) { // 감싸진 에러가 NumError인지 확인
    fmt.Println("NumberError:", numError)
}
```


---

## 23-II. 패닉

프로그램을 정상 진행시키기 어려운 상황에서 **프로그램 흐름을 중지**시키는 기능 

Go언어에서는 `panic()` 내장 함수 이용하여ㅓ 제공

### 23-II-1. 패닉 생성

```go
func painc(interface{})
```

panic() 내장 함수의 인수로 interface{} 타입 즉 모든 타입을 사용 가능

```go
painc(42)
panic("unreachable")
panic(fmt.Errorf("This is error num:%d", num))
panic(SomeType{SomeData})
```

특정 타입 객체 사용, fmt.Errorf() 함수 이용하여 만들어진 에러 타입 사용 가능


### 23-II-2. 패닉 전파 및 복구

에러 메시지 표시 및 복구 시도하는 것

- panic은 호출 순서를 **거슬러** 올라가며 전파됨
- 어느 단계에서든 panic은 **복구된 시점부터** 프로그램 계속됨
- `recover()` 함수 호출하여 패닉 복구 가능
    - 함수 호출 시점에 패닉 전파 중이면 panic 객체 반환, 그렇지 않으면 nil 반환

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("panic 복구 -", r)
    }
}()
```

### 23-II-3. recover() 결과

내장 함수 recover(): 발생한 panic 객체 반환

```go
func recover() interface{}
```

interface{} 타입 반환. 실제 사용 시 타입 검사 필요

```go
if r, ok := recover().(net.Error); ok {
    fmt.Println("r is net.Error Type")
}
```

위와 같이 발생한 패닉이 net.Error 타입인지 검사해서 처리.

발생한 패닉이 특정 타입인지 확인 후 해당 타입 패닉에 대한 처리 추가하고 싶은 경우 위와 같은 구문 사용하여 패닉 검사 시행하여야 함.

---

## 연습 문제

1. 

*os.PathError

2. 

숫자만 입력하세요. 문자:c

3.

```
15
nums should be not empty
end of main
```