# Chapter 31 Todo 리스트 웹 사이트 만들기

## 31.1 해법

1. 먼저 RESTful API에 맞춰서 서비스를 정의한다.
2. 할 일을 나타내는 Todo 구조체를 만든다.
3. 앞서 정의한 RESTful API에 맞춰서 각 핸들러를 만든다.
4. 화면을 구성하는 HTML 문서를 만든다.
5. 프론트엔드 동작을나타내는 자바스크립트 코드를 만든다.
6. 웹 브라우저로 동작을 확인한다.

## 31.2 준비하기

### 31.2.1 urfave/negroni 패키지 설치

urfave/negroni 패키지 : 자주 사용되는 웹 핸들러 제공하는 패키지

- 로그 기능 : 웹 요청을 받아 응답할 때 자동으로 로그를 남겨줘서 웹 서버 동작 확인 가능
- panic 복구 기능 : 웹 요청 수행 중 panic 발생 시 자동으로 복구해주는 동작 지원
- 파일 서버 기능 : public 폴더의 파일 서버 자동으로 지원

```go
m := MakeWebHandler()		//우리가 만든 핸들러
n := negroni.Classic() 	//negroni 기본 핸들러
n.UseHandler(m)					//우리가 만든 핸들러를 감쌈

err := http.ListenAndServe(":3000", n)	//negroni 기본 핸들러가 동작함
```

### 31.2.2 unrolled/render 패키지 설치

unrolled/render 패키지 : 웹 서버 응답을 구현하는 데 사용하는 유용한 패키지

```go
r := render.New()
r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
```

## 31.3 웹 서버 만들기

### 31.3.1 RESTful API 정의하기

| 메서드 | URL | 동작 |
| --- | --- | --- |
| GET | /todos | 전체 할 일 목록 반환 |
| POST | /todos | 새로운 할 일 등록 |
| PUT | /todos/id | id에 해당하는 할 일 업데이트 |
| DELETE | /todos/id | id에 해당하는 할 일 삭제 |

### 31.3.2 웹 서버 만들기

- 패키지 임포트, Todo 구조체 생성, 기타 전역 변수 선언
    - Todo 구조체 : ID, 이름, 완료 여부
    - `json` 은 JSON 포맷으로 변환 시 옵션

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type Todo struct { 
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

var todoMap map[int]Todo
var lastID int = 0
```

- 웹 서버 핸들러 생성

```go
func MakeWebHandler() http.Handler {
	rd = render.New()
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}
```

- Todos의 정렬 인터페이스 구현 & 할 일 목록 반환하는 핸들러 생성

```go
type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(w, http.StatusOK, list)
}
```

- 새로운 Todo 항목 등록

```go
func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}
```

- ID에 해당하는 Todo 항목 삭제, 삭제 성공 시 응답 & JSON 반환

```go
func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}
```

- ID에 해당하는 Todo 항목 수정

```go
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusBadRequest, Success{false})
	}
}
```

- main() 함수 구현

```go
func main() {
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
```

## 31.5 웹 배포 방법 고려하기

### 31.5.1 도메인이 없다

실제 서버 주소는 IP주소(구두점과 숫자로 이루어짐) → 도메인을 얻으려면 구매해야함

### 31.5.2 고정된 공개 IP 주소가 없다

일반적으로 개인이 고정된 공개 IP주소를 가지기 힘들다. 공유기나 인터넷 버시스 제공자가 컴퓨터에 할당한 임시 주소, 공유기 네트워크 안쪽에서만 유효하다.

호스팅 : 인터넷 호스팅 업체로부터 IP와 장비를 대여받아 웹 서비스 제공하는 서비스

→ 클라우드 서비스 : IaaS, PaaS, SaaS

## 연습문제

1. 

| 메서드 | URL | 동작 |
| --- | --- | --- |
| GET | /articles | 전체 게시물 반환 |
| POST | /articles | 새로운 게시물 등록 |
| PUT | /articles/id | id에 해당하는 게시물 업데이트 |
| DELETE | /articles/id | id에 해당하는 게시물 삭제 |
2. IaaS, PaaS, SaaS