# Chapter 4 RESTful API 서버 만들기

## 30.1 해법

1. gorilla/mux와 같은 RESTful API 웹 서버 제작을 도와주는 패키지를 설치한다.
2. RESTful API에 맞춰서 웹 핸들러 함수를 만들어준다.
3. RESTful API를 테스트하는 테스트 코드를 만든다.
4. 웹 브라우저로 데이터를 조회한다.

## 30.2 사전 지식 : RESTful API

REST : Representational State Transfer, 로이 필링이 2000년에 소개한 웹 아키텍처 설켸 형식

**HTTP 메서드**

| 메서드 | URL | 동작 |
| --- | --- | --- |
| GET | /students | 전체 학생 데이터 반환 |
| GET | /students/id | id에 해당하는 학생 데이터 반환 |
| POST | /students | 새로운 학생 등록 |
| PUT | /students/id | id에 해당하는 학생 데이터 변경 |
| DELETE | /students/id | id에 해당하는 학생 데이터 삭제 |

RESTful API 특징

1. 자기 표현적인 URL : URL만으로도 어떤 데이터에 대한 요청인지 알 수 있다.
2. 메서드로 행위 표현 : 메서드로 데이터에 대한 행위를 표현한다. URL과 메서드 조합으로 데이터에 대한 조작을 정의한다.
3. 서버/클라이언트 구조 : 서버는 데이터 제공자로 존재하고 클라이언트는 데이터 사용자로 동작한다.
4. 무상태 : 서버는 클라이언트의 상태를 유지하지 않는다. 서버가 상태를 보관할 필요가 없기 때문에 서버를 손쉽게 교체할 수 있어서 빠른 장애 대응이나 분산 처리에 유용하다.
5. 캐시 처리 : REST 구조로 서버가 단순해져서 더 쉽게 캐시 정책을 적용해서 성능을 개선할 수 있다.

## 30.3 RESTful API 서버 만들기

- 패키지를 임포트하고 전역 변수 선언 (students 학생 데이터 저장하는 맵 : Id가 키, Students 데이터 저장)

```go
package main

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int
```

- gorilla/mux 패키지를 이용해서 새로운 웹 핸들러를 만들고 임시 학생 데이터 두 개 생성해 저장
    - “/students” 요청을 받으면 GetStudentListHandler() 함수 호출
    - 임시 학생 데이터 두 개 생성해서 students에 저장
    - GetStudentListHandler() 호출 시 students 맵에 저장된 학생 데이터로 []Student 타입의 학생 목록 생성

```go

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}
```

- Id로 정렬하는 인터페이스 구현

```go
type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

```

- 학생 정보를 가져와 JSON 포맷으로 변경하는 핸들러 생성
    - Id로 정렬
    - JSON 포맷으로 변형

```go
func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

```

- main() 함수 : 3000번 포트에서 입력을 대기한다

```go
func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
```

## 30.4 테스트 코드 작성하기

- 테스트 코드 작성
    - “/students” 경로에 대한 GET 요청 생성
    - 웹 핸들러 요청 후 결과 반환 → []Student 타입의 list 변수를 만들어서 JSON 반환

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)                              
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
```

## 30.5 특정 학생 데이터 반환하기

- 메서드 “GET”으로 “/students/” 아래 숫자로 된 경로가 온다면 GetStudentHandler() 함수 호출하는 핸들러 등록

```go
mux.HandleFunc("/students/{id:[0-9]+}", GetStudentListHandler).Methods("GET")
```

- GetStudentHandler() 함수 추가
    - 자동으로 id값을 내부 맵에 저장
    - mux.Vars() 함수 호출하여 인수를 가져온다 → strconv.Atoi() 함수로 숫자로 변경
    - students 맵에서 id에 해당하는 학생 데이터가 있는지 확인
- GetStudentHandler() 테스트 코드

```go
func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
```

## 30.6 학생 데이터 추가/삭제하기

- 학생 데이터 추가 “/students” POST 요청

```go
mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
```

- PostStudentHandler() 함수 추가
    - 요청에 포함된 JSON 데이터를 Student 타입으로 변환
    - Id를 증가하여 맵에 등록한 뒤 응답코드 알림

```go
func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}
```

- 테스트 코드 추가

```go
func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`))

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}
```

- 학생 데이터 삭제 “/students/id” DELETE 요청

```go
mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")
```

- id에 해당하는 학생 데이터 삭제하는 DeleteStudentHandler() 함수 구현
    - id값 읽어오고 응답코드 반환

```go
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK)
}
```

- 테스트 코드 추가

```go
func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("bbb", list[0].Name)
}
```

## 연습문제

1. REST
2. id에 해당하는 뉴스 데이터 반환, id에 해당하는 뉴스 업데이트
3. 4, 5