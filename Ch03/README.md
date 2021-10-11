# Chapter 3  Hello Go World

### 3.1 Go 역사

2009년 발표된 오픈 소스 프로그래밍 언어

- 홈페이지 주소: [golang.org](http://golang.org)
- 온라인 Go 언어 컴파일러 : [play.golang.org](http://play.golang.org)

깃허브에 올라간 소스 코드 중 가장 많이 사용되는 언어 4위

### 3.2 Go 언어 특징

[Go 언어 특징](https://www.notion.so/b7450b3df3d741ef9768ac1f1ff32587)

### 3.3 코드가 실행되기까지

1. **폴더 생성** : 모든 코드는 패키지 단위, 같은 폴더 같은 패키지에 포함
2. **.go 파일 생성 및 작성** : 확장자는 반드시 .go로 끝나야함
3. **Go 모듈 생성** : 1.16 버전 이후 Go 모듈이 기본으로 적용, 빌드 전에 모듈 생성 (명령어 : **go mod init goproject/hello**), go.mod파일 생성됨
4. **빌드**  : Go 코드 → 기계어 (실행 파일), GOOS와 GOARCH 환경변수 조정으로 다른 OS, 아키텍처의 실행파일을 만들 수 있음, (빌드 가능한 OS, 아키텍처 목록 명령어 : go tool dist list)
5. **실행**

### 3.4 Hello Go World 코드 뜯어보기

```go
package main

import "fmt"

func main() {
  // Hello Go World 출력
	fmt.Println("Hello World")
}
```

1. **package main**
- 패키지 선언은 코드가 어떤 패키지에 속하는지 알려준다.
- main 패키지 : 프로그램 시작점(엔트리 포인트)을 포함하는 특별한 패키지
2. **import "fmt"**
- fmt 패키지 : 표준 입출력을 다루는 내장 패키지
3. **func main(){**
- main() 함수는 프로그램 진입점 함수
4. **// Hello Go World 출력**
- // : 한 줄 주석 예약어, /* */: 여러줄 주석
- 함수 앞에 함수명으로 시작하는 주석 → 함수 설명하도록
5. **fmt.Println("Hello Go World")**
- fmt.Println: 표준 출력 함수
6. **}**
- 코드 블록 종료

*****
### 느낀 점
C언어 같기도 Java 같기도...? 어떻게 하면 Go를 Go답게 쓸 수 있을지... 아직 많이 낯설지만 흥미롭다! 두근...