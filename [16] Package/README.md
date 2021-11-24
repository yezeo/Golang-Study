# Chapter 16. Package

## 16.1 패키지

패키지란, Go 언어에서 **코드를 묶는 가장 큰 단위** 이다.   

함수로 코드 블록을, 구조체로 데이터를 묶는다면,    

패키지로는 함수,구조체, 그외 코드를 묶을 수 있다.   

### 💡 main 패키지는 특별한 패키지

프로그램 시작점을 포함한 패키지이다 → 프로그램 시작점이란? main() 함수를 의미함.    

프로그램은 main 패키지(필수 요소) 하나와 여러 외부 패키지(선택 요소)로 구성된다.   

→ 프로그램 실행 과정을 알아보자,   

- 운영체제가 프로그램을 메모리로 올린다. ⇒ '로드(Load)'
- 이후, 프로그램 시작점(main함수)부터 한 줄씩 코드를 실행한다.

→ 이러한 main() 함수를 포함한 패키지가 main 패키지라는 걸 알아두자.

### 📌 그 외 패키지

- 표준 입출력 → fmt 패키지
- 암호화 기능 → crypto 패키지
- 네트워크 기능 → net 패키지

사용법은 import 하여 사용하면 된다.    

이미 세상에는 수많은 패키지가 제공되므로, 프로그램을 만들기 전에 이미 배포되어 있는 패키지를 먼저 찾아보면 도움이 된다.   

### 패키지 찾는 데 유용한 사이트

- [표준 패키지](https://pkg.go.dev/std)
- [Awesome Go](https://github.com/avelino/awesome-go)

## 16.2 패키지 사용법

```go
import (
	"fmt"
	"os"
)
```

- import 뒤에 패키지명을 큰따옴표로 묶어서 적는다.
- 소괄호로 패키지들을 묶어, 여러 패키지를 임포트 시킬 수 있다.

### 📌 패키지 멤버에 접근하기

```go
fmt.Println("Hello world")
```

위와 같이 패키지명을 쓰고 점 . 을 찍어서 패키지 내부의 함수를 호출하면 된다.

### 📌 경로가 있는 패키지 사용하기

```go
package main
import (
	"fmt"
	"math/rand"  // 패키지명은 rand 이다.
)
func main(){
	fmt.Println(rand.Int())  //랜덤한 숫자값을 출력하라
}
```

### 📌 겹치는 패키지 문제 → 별칭으로 해결

아래처럼 패키지명이 겹치는 경우,

```go
import (
	"text/template"  //template 패키지
	"html/template"  //template 패키지. 이름이 동일함
)
```

별칭 (aliasing)을 줘서 구별해줄 수 있다.

```go
import (
	"text/template"
	htemplate "html/template"  //별칭 htemplate 부여
)
```

### 📌 사용하지 않는 패키지 포함하기

패키지를 가져오면 반드시 사용해야 한다. 패키지를 임포트 하고 나서 사용하지 않으면, 에러가 발생한다.    

그러나 패키지를 직접 사용하지 않지만 부가효과를 얻고자 임포트 하는 경우가 있다.   

이런 경우에는 밑줄 _을 패키지명 앞에 붙여주어 해결할 수 있다.

```go
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"  //밑줄 _을 이용해서 오류 방지!
)
```

## 16.3 Go 모듈

Go 모듈은 Go 패키지들을 모아놓은 Go 프로젝트 단위이다.   

go build 를 하려면 반드시 Go 모듈 루트 폴더에 go.mod 파일이 있어야 한다.   

### 💡 go.mod 파일이란?

- 모듈 이름과 Go 버전, 필요한 외부 패키지 등이 명시되어 있는 파일.
- Go 언어에서는 go build 를 통해 실행 파일을 만들 때,
    - go.mod 파일과,
    - 외부 저장소 패키지 버전 정보를 담고 있는 go.sum 파일을 통해
    
    **외부 패키지**와 **모듈 내 패키지**를 합쳐서 **실행 파일**을 만들게 된다.    
    
- `go mod init [패키지명]` 명령어를 실행하여 go.mod 파일을 만들 수 있다.

## 🏁 Go 모듈을 만들고 외부 패키지 활용하기

1.    
goproject/usepkg 폴더에서   

 `go mod init goproject/usepkg` 명령으로 새로운 Go Module을 만들어주자.   

→ go.mod 파일 생성.   

1. 아래처럼 파일을 만들어주자.
- *goproject/usepkg/custompkg/custompkg.go*

```go
package custompkg //패키지 정의

import "fmt"

func PrintCustom() {
	fmt.Println("This is custom package!")
}
```

- *goproject/usepkg/usepkg.go*

```go
package main

import (
	"fmt"
	"goproject/ch16/usepkg/custompkg" //모듈 내 패키지 custompkg

	"github.com/guptarohit/asciigraph"           //외부 저장소 패키지 asciigraph
	"github.com/tuckersGo/musthaveGo/ch16/expkg" //외부 저장소 패키지 expkg
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
```

1. `go mod tidy` 실행   
- Go 모듈에 필요한 패키지를 찾아서 다운로드.
- 필요한 패키지 정보를 go.mod 파일과 go.sum 파일에 적어준다.

```bash
C:\Users\a\Desktop\goproject\ch16\usepkg>go mod tidy
go: finding module for package github.com/guptarohit/asciigraph
go: finding module for package github.com/tuckersGo/musthaveGo/ch16/expkg
go: downloading github.com/tuckersGo/musthaveGo/ch16/expkg v0.0.0-20210809125204-68bca0d80b54
go: downloading github.com/guptarohit/asciigraph v0.5.2
go: downloading github.com/tuckersGo/musthaveGo v1.0.0
go: found github.com/guptarohit/asciigraph in github.com/guptarohit/asciigraph v0.5.2
go: found github.com/tuckersGo/musthaveGo/ch16/expkg in github.com/tuckersGo/musthaveGo/ch16/expkg v0.0.0-20210809125204-68bca0d80b54
```

→ go.mod 파일이 변경되었고(아래 참고), go.sum 파일이 생성된 것을 확인할 수 있다.   

```go
module usepkg

go 1.17

require (
	github.com/guptarohit/asciigraph v0.5.2
	github.com/tuckersGo/musthaveGo/ch16/expkg v0.0.0-20210809125204-68bca0d80b54
)
```

1. `go build` 명령으로 실행 파일을 만들고, 실행해보자.
- 소소한 오류...
    
    [reference](https://codeac.tistory.com/127)
    
    ### 문제
    
    ```bash
    C:\Users\a\Desktop\goproject\ch16\usepkg>go build
    usepkg.go:5:2: package goproject/ch16/usepkg/custompkg is not in GOROOT 
    (C:\Program Files\Go\src\goproject\ch16\usepkg\custompkg)
    ```
    
    내가 의도한 `usepkg/custompkg`의 경로는 `goproject/ch16/usepkg/custompkg`인데,
    
    Go 컴파일러는 전혀 다른 경로인 `$GOROOT`에서 패키지를 찾고 있었다.
    
    저 경로가 아니면 패키지를 import할 수 없는지, 그럼 결국 매번 패키지를 만들면 `$GOROOT`에 넣어야만 하는지 고민하던 중, 의존성 관리 도구인 **Go Module**에 대해 제대로 알게 되었다.
    
    ### 해결
    
    `go mod init [패키지 경로]`을 통해 현재 디렉토리에 새 모듈을 생성한다.  
    
    즉, 1번에 적었던 go mod init으로 go.mod 파일을 생성할 때, 경로를 잘 적어줬어야 했다.    
    
    그래야 $GOROOT 경로가 잘 적용된다. 
    

→ 결과

```bash
C:\Users\a\Desktop\goproject\ch16\usepkg>.\usepkg.exe
This is custom package!
This is Github expkg Sample
 10.00 ┤        ╭╮
  9.00 ┤   ╭╮   ││
  8.00 ┤   ││ ╭╮││
  7.00 ┤   │╰╮││││╭╮
  6.00 ┤  ╭╯ │││││││ ╭
  5.00 ┤ ╭╯  ╰╯╰╯│││╭╯
  4.00 ┤╭╯       ││││
  3.00 ┼╯        ││││
  2.00 ┤         ╰╯╰╯
```

끝. 

> Go 모듈을 만들고 기본 패키지, 모듈 내부 패키지, 외부 저장소 패키지들을 조합하여 프로그램을 만들어 보았다. Go로 더 복잡한 응용 프로그램을 만들 때, 다른 언어들처럼 파일을 분리하고 임포트하여 구조적으로 사용할 수 있다는 걸 알아두자.
> 

### 💡 참고)

다운 받은 외부 패키지들은 GOPATH/pkg/mod 폴더에 버전별로 저장되어 있다.

→ 그래서 한 번 다운받은 패키지들은, 다른 모듈에서 사용될 때 다시 다운 받지 않고 사용하게 된다.

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/81405bac-c1ad-4c67-bd24-253e092e3904/Untitled.png)

## 16.4 패키지명과 패키지 외부 공개

패키지명은 쉽고 간단하게 이름 짓기, 모든 문자를 **소문자**로 할 것을 권장한다.   

📌 **외부 공개 조건**   

: 첫 글자가 **대문자**로 시작되는 모든 변수, 상수, 타입, 함수, 메서드, 구조체, 구조체의 필드는 패키지 외부로 공개.

## 🏁 외부 공개/비공개 알아보기

패키지 외부로 공개되는 것과 공개되지 않는 것을 예제로 알아보자.

- *goproject/ch16/ex16.2/publicpkg/publicpkg.go*

```go
package publicpkg

import "fmt"

const (
	PI = 3.1415   // 공개
	pi = 3.141516 //공개 x
)

var ScreenSize int = 1080
var screenHeight int

func PublicFunc() {
	const MyConst = 100
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() {
	fmt.Println("This is a private function")
}

type MyInt int
type myString string

type MyStruct struct {
	Age  int
	name string
}

func (m MyStruct) PublicMethod() {
	fmt.Println("This is a public method")
}

func (m MyStruct) privateMethod() {
	fmt.Println("This is a private method")
}

type myPrivateStruct struct {
	Age  int
	name string
}

func (m myPrivateStruct) PrivateMethod() {
	fmt.Println("This is a private method")
}
```

- *goproject/ch16/ex16.2/main.go*

```go
package main

import (
	"ch16/ex16.2/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
}
```

→ 결과

```bash
PI: 3.1415
This is a public function 100
myint: 10
mystruct: {18 }
```

- publicpkg의 pi 변수는 소문자로 시작하기 때문에 패키지 외부에서 접근할 수 없다.
- 역시 privateFunc() 함수도 외부로 공개되지 않는다. 따라서 패키지 외부에서 호출이 불가능.
- MyStruct구조체의 Age 필드는 대문자로 시작하기 때문에 접근할 수 있지만, name 필드는 접근 불가
