# Chapter 24 고루틴과 동시성 프로그래밍

## 24.1 스레드

고루틴: 경량 스레드, 함수나 명령을 동시에 실행할 때 사용, main() 함수도 고루틴에 의해서 실행

프로세스: 메모리 공간에 로딩되어 동작하는 프로그램, 스레드를 한 개 이상 가짐

스레드: 프로세스 안의 세부 작업 단위, 실행 흐름

### 24.1.1 컨텍스트 스위칭 비용

컨텍스트 스위칭 비용: CPU 코어가 여러 스레드를 전환하면서 수행하면 더 많은 비용이 든다

스레드를 전환하려면 현재 상태를 보관해야함 ⇒ 스레드의 명령 포인터

스택 메모리 등의 정보 저장 ⇒ 스레드 컨텍스트

적정 개수를 넘어 한 번에 너무 많은 스레드 수행 → 성능 저하

하지만 Go 언어에서는 CPU 코어마다 OS 스레드를 하나만 할당해서 사용 (컨텍스트 스위칭 비용 발생 X)

## 24.2 고루틴 사용

모든 프로그램은 고루틴을 최소 하나 가짐 (= 메인 루틴)

```go
	go 함수_호출 //고루틴 추가 생성
```

```go
package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
}
```

출력문

```go
가 1 나 2 다 라 3 마 4 바 5 사
```

- PrintHangul()과 PrintNumbers() 함수는 각기 다른 새로운 고루틴에서 실행 (동시에 실행)
- 메인루틴에서 3초간 대기 : 메인 함수가 종료되면 아무리 많은 고루틴이 있더라도 프로그램 즉시 종료 (아무 출력도 없음)

### 24.2.1 서브 고루틴이 종료될 때까지 기다리기

sync 패키지의 WaitGroup 객체 사용

```go
var wg sync.WaitGroup

wg.Add(3)   //작업 개수 설정
wg.Done()   //작업 완료될 때마다 호출
wg.Wait()   //모든 작업이 완료될 때까지 대기
```

## 24.3 고루틴의 동작 방법

## 24.4 동시성 프로그래밍 주의점

동시성 프로그래밍의 문제점: 동일한 메모리 자원에 여러 고루틴이 접근할 때 발생

고루틴은 각 CPU 코어에서 별도로 동작하지만, 고루틴은 같은 메모리 공간에 동시에 접근해서 값 변경 가능

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 { //잔고가 0이하면 패닉
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
```

출력문

```go
panic: Balance should not be negative value: -2000

goroutine 15 [running]:
main.DepositAndWithdraw(0xc0000140b0)
...
```

- 여러 고루틴에서 통장에 동시 접근 → 동시성 문제 발생
- DepositAndWithdraw() 함수 무한히 호출 → 시간이 지나면 잔고가 0이 되면서 패닉 발생

## 24.5 뮤텍스를 이용한 동시성 문제 해결

고루틴에서 값을 변경할 때 다른 고루틴이 건들지 못하게 하는 것 = 뮤텍스 이용하여 자원 접근 권한 통제

뮤텍스의 Lock() 메서드 호출 → 획득 or 대기, Unlock() 메서드 호출 → 반납

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```

- 출력이나 패닉 없이 무한 대기
- DepositAndWithdraw() 함수에서 mutex.Lock() 메서드 호출 : 뮤텍스 획득
- defer mutex.Unlock() 메서드 : 함수 종료 전에 반납
- 잔고는 절대 0원 이하로 내려가지 않게 된다

## 24.6 뮤텍스와 데드락

뮤텍스로 동시성 프로그래밍 문제 해결 가능, 또 다른 문제!

1. 동시성 프로그래밍으로 얻는 성능 향상을 얻을 수 없다: 여러 고루팅 중 뮤텍스를 획득한 고루틴만 실행
2. 데드락이 발생할 수 있다

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
```

출력문

```go
A 밥을 먹으려 합니다.
A 포크 획득
B 밥을 먹으려 합니다.
A 수저 획득
A 밥을 먹습니다
A 밥을 먹으려 합니다.
A 포크 획득
B 수저 획득
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
...
```

- 포크와 수저 뮤텍스, A는 포크 먼저, B는 수저 먼저
- diningProblem() 함수 : 첫번째, 두번째 뮤텍스 모두 획득 → 밥 → 반납
- A, B 서로 두번째 뮤텍스 획득X 무한 대기 → 데드락 감지 & 에러 반환

실제 프로그래밍에서 뮤텍스들이 복잡하게 꼬여 있어서 단순히 순서 변경으로 해결할 수 없는 경우가 있다. 데드락 문제는 동시성 프로그래밍에서 해결하기 힘든 난제

정리 

- 멀티코어 컴퓨터에서 여러 고루틴 사용하여 성능 향상 가능
- 같은 메모리를 여러 고루틴이 접근하면 프로그램이 꼬일 수 있다
- 뮤텍스 이용하여 동시에 고루틴 하나만 접근하도록 조정
- 뮤텍스를 잘못 사용하면 성능 향상 X, 데드락 문제 발생 가능

## 24.7 또 다른 자원 관리 기법

모든 문제는 같은 자원을 여러 고루틴이 접근하기 때문에 발생한다. 멀티코어의 이점 + 뮤텍스 X ⇒ 각 고루틴이 서로 다른 자원에 접근하게 만드는 두 가지 방법

- 영역을 나누는 방법
- 역할을 나누는 방법

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
```

출력문

```go
9 작업 시작
4 작업 시작
0 작업 시작
1 작업 시작
2 작업 시작
3 작업 시작
6 작업 시작
5 작업 시작
7 작업 시작
8 작업 시작
6 작업 완료 - 결과: 36
2 작업 완료 - 결과: 4
1 작업 완료 - 결과: 1
4 작업 완료 - 결과: 16
9 작업 완료 - 결과: 81
0 작업 완료 - 결과: 0
3 작업 완료 - 결과: 9
5 작업 완료 - 결과: 25
8 작업 완료 - 결과: 64
7 작업 완료 - 결과: 49
```

## 연습문제

1. 1, 2, 3
2. mutex.Unlock(), wg.Done(), 10

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))	
	}
	account.Balance += 1000
	account.Balance -= 1000
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	account := Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			DepositAndWithdraw(account)
		}()
	}
	wg.Wait()
}
```