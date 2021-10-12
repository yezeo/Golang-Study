#### 목차
[1- Go 언어 소개](https://www.notion.so/2-Hello-Go-World-7d2b788882434cf099d693be63804e6c)  
[2- Go 언어의 특징](https://www.notion.so/2-Hello-Go-World-7d2b788882434cf099d693be63804e6c)  
[3- 코드 뜯어보기](https://www.notion.so/2-Hello-Go-World-7d2b788882434cf099d693be63804e6c)
 
## 2.1 Go 소개

**📌 역사**

- 2009년 구글에서 만든 오픈소스 프로그래밍 언어이다.
    
    → 누구나 내부 구조를 살펴볼 수 있고, 언어 발전에 이바지 가능.
    
- Go, Golang, 고랭 등 다양하게 불린다.

**📌 장점**

- 심플한 문법 구조를 가지고 있다.
    
    ⇒ **적은 코딩**으로 **빠르고 강력한 성능**을 낼 수 있다.
    
- 개발자가 사랑하는 언어 5위.
- 미국에서 3번째로 많은 연봉을 받는 인기 언어.

**📌 어디에 쓰이는가**

- 백엔드 서버와 시스템 개발에 적합하고 강력한 동시성 프로그래밍을 지원한다.

## 2.2 Go 언어의 특징

### ( **O )**

- **메서드** →
- **인터페이스** →
- **익명 함수** →
- **가비지 컬렉터** →
- **포인터** →

### ( X )

- **클래스** →
- **상속** →
- **제네릭 프로그래밍** →
- **네임스페이스** → 모든 코드는 패키지 단위로 분리된다.

---

## 2.3 Go 코딩

코드 실행 시 Go는 이런 과정을 거칩니다.

**`Go 모듈 생성`** **> `빌드`** (.exe 파일 생성) **>** **`실행`**

### 📌 Go 파일 실행 방법 (Terminal)

0. (vscode) Terminal > New Terminal 에서 아래 명령어 입력

1. `cd (폴더명)`
2. `go mod init 프로젝트명/폴더명`  //모듈 생성
3. `go build`  //실행하면 아래와 같이 .exe 파일이 생성된다.

![Untitled 6](https://user-images.githubusercontent.com/61778930/136964413-f011cf1a-aae9-48e5-a872-8d7678d71f33.png)


4. `.\hello.exe`  // .exe파일 실행

 → **성공!**

![Untitled 7](https://user-images.githubusercontent.com/61778930/136964459-2ca7b6c0-6a29-4e78-ab81-d4033c43493e.png)


---

### 📌 실행 원리 알아보기

1) `cd hello`

2) `go mod init goproject/hello` 명령으로 모듈을 생성한다.

모든 Go 코드는 빌드하기 전에 **모듈을 생성**해야 한다!

→ go.mod 파일이 goproject/hello 폴더 내에 생성됨.

→ go.mod 파일에는 모듈명, Go 버전, 필요한 패키지 목록 정보가 담겨 있음.

3) `go build` 로 빌드

Go 코드를 **기계어로 변환**하여 실행 파일을 만든다.

+ 이때, 실행파일에 대한 설정을 GOOS와 GOARCH 환경변수를 이용해 바꿀 수 있다.

ex) `GOOS=linux GOARCH=amd64 go build` → AMD64 계열의 리눅스 실행 파일 만들기

4) `.\hello.exe` 로 실행

---

### 📌 Hello Go World 코드 뜯어보기

코드 한 줄 한 줄이 의미하는 바를 알아보자.

```go
package main   //1

import "fmt"   //2

func main() {  //3
	fmt.Println("Hello Go World!")   //4
}
```

1) `package main` → Go언어의 모든 코드는 반드시 패키지 선언으로 시작해야 한다. 

main 패키지에 속한 코드라고 컴파일러에게 알려준다.

main 패키지는 "**프로젝트 시작점이 있는 패키지다**."

⇒ 무슨 말이냐? main() 함수가 없는 패키지는 패키지 이름으로 main을 쓸 수 없다. 
프로그램은 항상 main() 함수에서 시작되고, main() 함수가 종료되면 프로그램도 종료된다. 
즉, 프로그램 시작과 끝이 main() 함수이므로,

 패키지 이름으로 main을 쓰는 **package main**은 프로그램 시작점이 있는 패키지이다.

2) `import "fmt"` → 표준입출력을 다루는 **fmt** 패키지를 불러온다.

3) `func main() { }` → main() 함수는 프로그램 진입점 함수!

4) `fmt.Println("Hello Go World")` → Hello Go World를 터미널 화면에 출력합니다. (표준 출력)
