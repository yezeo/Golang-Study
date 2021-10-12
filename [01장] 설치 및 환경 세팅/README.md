## 설치

[https://golang.org/dl/](https://golang.org/dl/) 에 접속해서 윈도우용 파일 다운로드,
   
![Untitled](https://user-images.githubusercontent.com/61778930/136963793-86a88466-5c6f-4175-bdc1-1c2a7447593f.png)

다운로드한 설치 파일을 클릭하여 실행한다. > 이후 설치 진행

![Untitled 1](https://user-images.githubusercontent.com/61778930/136963363-0172f551-e48b-4b47-85b3-5dea200ea335.png)


## 설치 확인
![Untitled 2](https://user-images.githubusercontent.com/61778930/136963502-f949881b-7006-48b0-a525-eea6bcf99574.png)

## Go 확장 프로그램 설치

IDE는 VScode를 사용한다. 

vscode에서 Go 프로그래밍을 할 수 있도록 도와주는 확장 프로그램을 설치해보자.

vscode에서 상단 메뉴 View>Extensions에 들어가, 
검색창에 'Go' 입력 후

해당 익스텐션을 다운 받는다.

![Untitled 3](https://user-images.githubusercontent.com/61778930/136963649-24ee507c-4056-4c03-a838-a1e246b28793.png)

vscode 우측 하단에 뜨는 팝업창 > Install All 클릭.

output 창에 아래와 같은 문구가 뜨면 모든 환경 구축 완료.

![Untitled 4](https://user-images.githubusercontent.com/61778930/136963687-b7422658-8509-47fd-8001-e53e0530dd08.png)


---

## 코드 실행해보기

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go World!")
		//fmt.println("Hello Go World!")
}
```

간단히 "hello go world"를 출력하는 첫 코드를 작성하고 실행했는데, 

`fmt.println` 함수를 찾을 수 없다는 에러가 떴다. 

왜 안 될까, 설정을 잘못 했나 하고 패키지도 옮겨보고 구글링해서 캐시도 지우고 했지만 

알고보니 내가 함수명을 잘못쓴 거였다. 

`fmt.Println` 처럼 함수가 대문자로 시작해야 했다. 허무했다.

뭐든 처음이 어려운 것 같다 ^_^;;; 얼른 익숙해지고 싶다.

![Untitled 5](https://user-images.githubusercontent.com/61778930/136963714-8a024f13-360e-4a00-9ad0-b46ca8d947c8.png)


> Notice how we have written fmt.println and not `fmt.Println`. 
Go is case dependent, which means to say that when you use another’s name, 
use it exactly as it is defined. 
If the name is John, then only **John** works - not john, not joHn, and no other combination.
> 

^^..

---

## Go 파일 실행 순서

1. `cd (패키지명)`
2. `go mod init 프로젝트명/패키지명`
3. `go build`  //실행하면 아래와 같이 .exe 파일이 생성된다.


4. `.\hello.exe`  // exe파일 실행

![Untitled 6](https://user-images.githubusercontent.com/61778930/136963736-206bca2b-c07f-447c-9965-3a639cf0af17.png)

 → **성공!**
    
![Untitled 7](https://user-images.githubusercontent.com/61778930/136963773-0757ec13-2926-4bfb-b90a-a35c917e083a.png)


---

[https://litaro.tistory.com/entry/Go-언어-초보의-Go-modules-정리-노트-1](https://litaro.tistory.com/entry/Go-%EC%96%B8%EC%96%B4-%EC%B4%88%EB%B3%B4%EC%9D%98-Go-modules-%EC%A0%95%EB%A6%AC-%EB%85%B8%ED%8A%B8-1)
