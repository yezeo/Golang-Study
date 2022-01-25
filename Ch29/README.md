# Chapter 29 Go 언어로 만드는 웹 서버

## 29.1 HTTP 웹 서버 만들기

### 29.1.1 핸들러 등록

핸들러 : 각 HTTP 요청 URL이 수신됐을 때 그것을 처리하는 함수 또는 객체

HandleFunc() 함수로 핸들러 등록

Handle() 함수로 http.Handler 인터페이스를 구현한 객체 등록 

URL 경로에 해당하는 HTTP 요청 수신 시 핸들러에 해당하는 함수 호출 또는 http.Handler 객체의 인터페이스인 ServeHTTP() 헤서드 호출 → 요청에 따른 로직 수행

```go
func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	...
}

http.HandleFunc("/", IndexPathHandler)
```

**http.Request 구조체 살펴보기**

```go
type Request struct {
	Method string //메서드 정보
	URL *url.URL

	Proto      string //프로토콜 버전 정보
	ProtoMajor int
	ProtoMinor int

	Header header
	Body io.ReadCloser
	...
}
```

### 29.1.2 웹 서버 시작

웹 서버 시작

```go
func ListenAndServe(addr string, handler Handler) error
```

addr : HTTP 요청을 수신하는 주소, 포트번호

handler : 핸들러 인스턴스, nil로 넣으면 디폴트 핸들러 실행

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.ListenAndServe(":3000", nil)
}
```

- HandleFunc() → 웹 핸들러 함수 등록
- 첫번째 인수 : “/” 루트 경로
- 두번째 인수 : 실행할 함수
    - 첫번째 인수 : http.ResponseWriter
    - 두번째 인수 : *http.Request (메서드, 헤더, 바디 등 HTTP 요청 정보)
- fmt.Fprint() : 지정한 출력 스트림에 출력한다 (http.ResponseWriter 타입을 출력 스트림으로 지정)
- ListenAndServe() 함수 : 웹 서버 실행
    - 첫번째 인수 : 요청을 기다릴 주소 (3000번 포트)
    - 핸들러 인스턴스 : nil은 DefaultServeMux 사용

## 29.3 HTTP 쿼리 인수 사용하기

쿼리 인수 : URL 끝에 붙여넣는 인수

- http://localhost?id=1&name=abcd

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
```

- Query() 메서드 : 쿼리 인수 가져오기, 결괏값 = Value 타입 (map[string[]string 타입의 별칭)
- values.Get(”name”) : name 키값의 인숫값 가져오기 → strconv.Atoi() 함수로 int 타입으로 변경
- “/bar” 경로에 핸들러 등록

## 29.4 ServeMux 인스턴스 이용하기

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListenAndServe(":3000", mux)
}
```

- http.NewServeMux() 함수 : 새로운 ServeMux 인스턴스 생성
- http.HandleFunc() 함수는 DefaultServeMux에 핸들러 등록 vs mux.HandleFunc() 메서드는 ServeMux 인스턴스에 핸들러 등록

**Mux**

multiplexer의 약자로 여러 입력 중 하나를 선택해서 반환하는 디지털 장치 (웹 서버는 각 URL에 해당하는 핸들러 등록 →  요청이 오면 해당 핸들러 선택 & 실행)

## 29.5 파일 서버

```html
<!--ch29/ex29.4/test.html-->
<html>
<body>
<img src="https://golang.org/lib/godoc/images/footer-gopher.jpg"/>
<h1>이것은 Gopher 이미지입니다.</h1>
</body>
</html>
```

<img> 태그의 src값으로 이미지 경로를 가짐

웹브라우저는 HTML을 표시하는 데 필요한 이미지 데이터를 다시 HTTP 요청을 통해서 가져오게 된다. 이미지 요청을 받은 웹 서버는 이미지 경로에 해당하는 데이터를 반환해서 화면에 이미지 표시할 수 있도록 한다.

### 29.5.1 “/” 경로에 있는 파일 읽어오기

```go
package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", nil)
}
```

### 29.5.2 특정 경로에 있는 파일 읽어오기

```go
http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
```

- [http://localhost:3000/static/gopher.jpg](http://localhost:3000/static/gopher.jpg) 경로의 이미지 데이터 요청 → 웹 서버는 해당 경로의 이미지 데이터를 웹 브라우저로 반환 → 이미지 출력

## 29.6 웹 서버 테스트 코드 만들기

```go
package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
```

- MakeServeMux() 함수에서 핸들러 인스턴스 생성

```go
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // / 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) 
	data, _ := io.ReadAll(res.Body)       
	assert.Equal("Hello World", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // /bar 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
```

- httptest 패키지의 NewRequest() 함수 : 테스트용 경로 요청 객체 생성
- MakeServeMux() 함수 : ex29.5.go 파일에서 만든 핸들러 인스턴스를 가져와 테스트 → ServeHTTP() 메서드 호출
- http.StatusOK 확인 (문제 없음)
- io.ReadAll() 함수 : Request 객체의 Body 읽어오기 ([]byte 타입 → string 타입)

## 29.7 JSON 데이터 전송

JSON : 자바스크립트 오브젝트 표기법

표기 규칙

- 오브젝트 시작은 { 로 표기, } 로 종료
- 필드는 “key” : value 형태로 표기
- 각 필드는 , 로 구분
- 배열은 [ ]로 표기
- 문자열은 “ “ 로 묶어서 표기

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler { 
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) 
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
```

- MakeWebHandler() 함수 : 핸들러 인스턴스 반환
- Student 객체 → JSON 포맷 ([]byte 타입)

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) // /student 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)
	assert.Equal(16, student.Age)
	assert.Equal(87, student.Score)
}
```

- json.NewDecoder(res.Body).Decode(student) : JSON 데이터 → Student 객체 → 값 확인

## 29.8 HTTPS 웹 서버 만들기

HTTPS : HTTP에 보안 기능을 강화한 프로토콜, 기존 HTTP 요청과 응답을 공개키 암호화 암호화 방식을 사용해 암호화

### 29.8.1 공개키 암호화 방식

공개키 암호화 방식 

- 공개키와 비밀키 두 가지 키를 생성
- 공개키는 클라리언트에 알려주고 비밀키는 서버에 비공개 상태로
- 클라이언트에서 HTTPS 요청을 보낼 때 공개키로 암호화, 서버는 비밀키를 이용해 복호화
- 비대칭 암호화 방식 (암호화와 복호화에 쓰이는 키가 서로 다름)

### 29.8.2 인증서와 키 생성

openssl 프로그램로 인증서 발급

```go
openssl req -new -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr
```

rsa:2048 방식으로 키 생성 → 비밀키는 localhost.key 파일로 저장 & 인증 파일은 localhost.csr로 저장 → 인증 기간과 제출해서 인증서인 .crt 파일 생성

셀프 인증

```go
openssl x509 -req -days 365-in localhost.csr -signkey localhost.key -out localhost.crt
```

 대표적인 인증 알고리즘 x509로 인증서 발급 → localhost.crt 파일 생성 (공개키 포함)

서버 실행 & 클라이언트 접속 시 클라이언트에 localhost.crt 파일로 인증 정보와 공개키 전송

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil) 
	if err != nil {
		log.Fatal(err)
	}
}
```

## 연습문제

1. 4
2. “net/http”, ResponseWriter