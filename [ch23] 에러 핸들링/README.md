# Chapter 23. Error Handling
## 23.1 에러 반환

기본 처리: 에러 반환 후 알맞게 처리

ex) 
ReadFile() 함수 → 해당 파일 없는 에러 발생 시, 

메세지 출력 후 다른 파일 읽거나 임시 파일 생성

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) 
	if err != nil {
		return "", err
	}
	defer file.Close()         
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) 
	if err != nil {                  
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) 
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) 
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") 
		if err != nil {                                
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename) 
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line)
}
```

결과

```go
파일내용: This is WriteFile
```

### ReadFile() 함수

- os.Open() 함수로 파일 열기 → 두 번째 반환값 error 실패 : 에러 발생, file.Close() 함수 defer 키워드 사용하여 파일 핸들 닫기
- bufio.NewReader() 함수로 객체 생성 : ‘\n’ 나올 때까지 파일에서 문자열 읽기, 두번째 반환값 error는 문자열이 delim 문자로 끝나지 않을 경우만 (밑줄 _ 로 에러 무시)

### WriteFile() 함수

- os.Create() 함수로 파일 생성
- fmt.Fprintln() 함수로 파일 핸들에 문자열 작성: 첫번째 인수로 Write() 메서드를 가진 io.Writer 인터페이스, 두번째 인수로 쓸 내용, 첫번째 반환값으로 쓴 길이 (밑줄 _로 무시), 두번째 반환값으로 에러 발생 시 에러

### main() 함수

- ReadFile() 함수 시도 → 에러 발생 시 WriteFile() 함수 호출로 파일 생성 → 에러 발생 시 에러 메세지 출력 후 종료 / 성공 시 ReadFile() 함수 시도 후 내용 출력

### 📌 **사용자 에러 반환**

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(
			"제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
```

결과:  `Error: 제곱근은 양수여야 합니다. f:-2`

- Sqrt() 함수 : 인수의 제곱근 반환, 인수 f가 음수이면 에러 반환
- fmt.Errorf() 함수로 에러 반환 → 원하는 에러 메시지 생성 가능
- errors 패키지의 New() 함수로 error 생성 가능

## 23.3 패닉
