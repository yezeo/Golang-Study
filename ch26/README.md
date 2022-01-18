# Chapter 26 Project 단어 검색 프로그램 만들기

## 26.1 해법

1. 경로에 해당하는 파일을 찾는다. 파일 경로는 특정 파일 하나만 나타낼 수도 있고 (와일드 카드로 표현되어) 여러 파일을 나타낼 수도 있다. 
2. 찾으려는 단어와 파일 경로를 입력 받는다. 프로그램은 입력받은 인수를 읽어서 사용해야 한다.
3. 파일을 읽고 각 라인에서 해당 단어가 나오는지 확인한다.
4. 특정 단어가 등장하는 라인을 취합하여 마지막으로 결과를 출력해야 한다.

![0](.\img\0.png)

## 26.2 사전 지식

### 26.2.1 와일드카드

파일 경로를 표시할 때 와일드카드를 사용해서 여러 파일을 나타낼 수 있다.

| * | 0개 이상의 아무 문자 | 예) abc*t ⇒ abct, abc34t, abcffdt ... |
| --- | --- | --- |
| ? | 1개의 아무 문자 | 예) ab?d ⇒ abcd, ab2d, abtd ... |

### 26.2.2 os.Args 변수와 실행 인수

보통 터미널 명령을 실행할 때 실행 인수를 넣어서 명령의 행동을 조정한다.

Go 언어에서는 os 패키지의 Args 변수를 이용해 실행 인수를 가져올 수 있다.

Args는 os 패키지의 전역 변수로 각 실행 인수가 []string 슬라이스에 담겨 있다.

```go
var Args []string
```

os.Args의 첫 번째 항목 : 실행 명령, 두 번째 항목부터 : 입력한 인수

```go
find word filepath
```

os.Args[0] : find

os.Args[1] : 찾을 단어

os.Args[2] : 파일 경로

실행 인수를 가져온 다음 파일 경로에 해당하는 파일 목록을 가져와야 한다. 파일 경로는 하나의 파일을 지정할 수 있지만 ? 혹은 * 같은 와일드카드를 사용해서 여러 파일을 지정할 수도 있다.

### 26.2.3 파일 핸들링

**파일 열기**

os 패키지의 Open() 함수 : 파일 열어 파일 핸들 가져오기

name에 해당하는 파일을 읽기 전용으로 읽고 *File 타입인 파일 핸들 객체 반환

```go
func Open(name string) (*File, error)
```

*File 타입은 io.Reader 인터페이스를 구현하고 있기 때문에 bufio 패키지의 NewScanner() 함수를 통해 스캐너 객체를 만들어서 사용할 수 있다.

**파일 목록 가져오기**

path/filepath 패키지의 Glob()함수 : 파일 경로에 해당하는 파일 목록 가져오기

```go
func Glob(pattern string) (matches []string, err error)
```

파일 경로를 넣어주면 경로에 해당하는 파일 리스트를 []string타입으로 반환

```go
filepaths, err := filepath.Glob("*.txt")
```

현재 실행 폴더 내 확장자가 txt인 모든 파일 리스트 반환

**파일 내용 한 줄씩 읽기**

bufio 패키지의 NewScanner() 함수 : 파일을 한 줄씩 읽기

```go
func NewScanner(r io.Reader) *Scanner
```

io.Reader 인터페이스를 구현한 모든 인스턴스를 인수로 사용 가능

*File 객체 또한 io.Reader 인터페이스를 구현하고 있어서 앞서 구한 파일 핸들을 사용해서 스캐너를 만들 수 있다. 스캐너를 이용하면 파일을 한 줄씩 손쉽게 읽어올 수 있다.

```go
type Scanner
	func (s *Scanner) Scan() bool
	func (s *Scanner) Text() string
```

Scan() 메서드 : 다음 줄 읽어오기

Text() 메서드 : 읽어온 한 줄을 문자열로 반환하기

### 26.2.4 단어 포함 여부 검사

strings 패키지의 Contains() 함수 : 읽어온 한 줄 내용 중에 찾으려는 단어가 있는지 검사하기

```go
func Contains(s, substr string) bool
```

strings.Contains() : 첫 번째 인수 s 안에 두 번째 인수 substr이 포함되어 있는지 여부를 반환

## 26.3 (step 1) 실행 인수 읽고 파일 목록 가져오기

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 { //실행 인수 개수
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1] //실행 인수 가져오기
	files := os.Args[2:]
	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path) //파일목록 가져오기
		if err != nil {
			fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
			return
		}
		fmt.Println("찾으려는 파일 리스트")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
```

출력문

```go
C:\Users\WINDOWS10\Desktop\Golang-Study\ch26\ex26.1>ex26.1 word ex*
찾으려는 단어: word
찾으려는 파일 리스트
ex26.1.exe
ex26.1.go
```

## 26.4 (step 2) 파일을 열어서 라인 읽기

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}
func PrintFile(filename string) {
	file, err := os.Open(filename) //파일 열기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return
	}
	defer file.Close() //함수 종료 전에 파일 닫기

	scanner := bufio.NewScanner(file) //스캐너를 생성해서 한 줄씩 읽기
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```

출력문

```go
THE TRAGEDY OF HAMLET, PRINCE OF DENMARK

by William Shakespeare
...
THE END
```

## 26.5 (step 3) 파일 검색 프로그램 완성하기

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//찾은 라인 정보
type LineInfo struct {
	lineNo int
	line   string
}

//파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.3 word filepath")
		return
	}

	word := os.Args[1] //찾으려는 단어
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		//파일 찾기
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------")
		fmt.Println()
	}
}
```

- 단어가 포함된 한 줄 텍스트 정보를 포함하는 구조체 선언
- 실행 인수에서 찾으려는 단어와 파일명 가져오기
- FindWordInAllFiles() 함수 호출해서 파일 내 단어 찾기 → 반환 []FindInfo를 이용해서 탐색 결과 출력

```go
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path) //파일 리스트 가져오기
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	for _, filename := range filelist { //각 파일별로 검색
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file) //스캐너 생성
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) { //한 줄씩 읽으면 단어 포함 여부 검색
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}
```

- filepath의 Glob() 함수를 통해 path에 해당하는 파일 리스트 가져오기
- 파일별로 FindWordInFile() 함수 실행해서 파일 검사 (파일 열고 각 줄을 읽은 뒤 원하는 단어가 있는지 검사, 있을 경우 결과 리스트에 추가)

출력문

```go
C:\Users\WINDOWS10\Desktop\Golang-Study\ch26\ex26.3>ex26.3 brother *.txt
hamlet.txt
--------------------------------
         267       King. Though yet of Hamlet our dear brother's death
         285         Or thinking by our late dear brother's death
         291         To our most valiant brother. So much for him.
         422         My father's brother, but no more 
like my father
         606         As watchman to my heart. But, good my brother,
         908         Thus was I, sleeping, by a brother's hand
         1269        Say, Voltemand, what from our brother Norway?
         2539        A brother's murther! Pray can I not,
         2545        Were thicker than itself with brother's blood,
         2632        You are the Queen, your husband's brother's wife,
         2648        As kill a king, and marry with his brother.
         2677        The counterfeit presentment of two brothers.
         2688        Blasting his wholesome brother. Have you eyes?
         3234        My brother shall know of it; and 
so I thank you for your good
         3251        Her brother is in secret come from France;
         3975      Ham. I lov'd Ophelia. Forty thousand brothers
         4258        And hurt my brother.
         4269        And will this brother's wager frankly play.
--------------------------------

romeonjuliet.txt
--------------------------------
         636           Mercutio and his brother Valentine;
         2419      Cap. Wife. Tybalt, my cousin! O my 
brother's child!
         3149        But for the sunset of my brother's son
         4030      John. Holy Franciscan friar, brother, ho!
         4039      John. Going to find a barefoot brother out,
         4054      Laur. Unhappy fortune! By my brotherhood,
         4473      Cap. O brother Montague, give me thy hand.
--------------------------------
```

## 26.6 개선하기

ex26.2 구현은 모든 파일 검색을 하나의 main() 고루틴에서 실행하지만 파일 개수가 늘어나면 검색에 오랜 시간이 든다. → 고루틴을 사용해서 파일 개수가 늘어도 빠르게 검색되도록 개선하기!

각 파일별로 작업을 할당하고 작업이 완료되면 채널을 이용해서 결과를 수집하는 방식 사용하기

![1](.\img\1.png)

뿌리고 거두기 방식

```go
func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	filelist, err := filepath.Glob(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch) //고루틴 실행
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo) //결과 수집
		recvCnt++
		if recvCnt == cnt {
			// all received
			break
		}
	}
	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo			//채널에 결과 전송
}
```

- FindWordInAllFiles() 함수에서 파일별로 고루틴 생성 & FindWordInFile() 함수 실행
- 채널을 통해서 결과 수집 → 모든 결과 수집되면 호출자에게 결과 알리기

출력문

```go
C:\Users\WINDOWS10\Desktop\Golang-Study\ch26\ex26.4>ex26.4 brother *.txt
hamlet.txt
--------------------------------
         267       King. Though yet of Hamlet our dear brother's death
         285         Or thinking by our late dear brother's death
         291         To our most valiant brother. So much for him.
         422         My father's brother, but no more 
like my father
         606         As watchman to my heart. But, good my brother,
         908         Thus was I, sleeping, by a brother's hand
         1269        Say, Voltemand, what from our brother Norway?
         2539        A brother's murther! Pray can I not,
         2545        Were thicker than itself with brother's blood,
         2632        You are the Queen, your husband's brother's wife,
         2648        As kill a king, and marry with his brother.
         2677        The counterfeit presentment of two brothers.
         2688        Blasting his wholesome brother. Have you eyes?
         3234        My brother shall know of it; and 
so I thank you for your good
         3251        Her brother is in secret come from France;
         3975      Ham. I lov'd Ophelia. Forty thousand brothers
         4258        And hurt my brother.
         4269        And will this brother's wager frankly play.
--------------------------------

romeonjuliet.txt
--------------------------------
         636           Mercutio and his brother Valentine;
         2419      Cap. Wife. Tybalt, my cousin! O my 
brother's child!
         3149        But for the sunset of my brother's son
         4030      John. Holy Franciscan friar, brother, ho!
         4039      John. Going to find a barefoot brother out,
         4054      Laur. Unhappy fortune! By my brotherhood,
         4473      Cap. O brother Montague, give me thy hand.
--------------------------------
```

## 연습문제

```go
func GetFileList(pattern string) ([]string, error) {
	filelist := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				filelist = append(filelist, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return filelist, nil
}
```