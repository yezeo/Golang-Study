# Chapter 24. goroutine

## 24.1 스레드란?

**고루틴**: 경량 스레드

- 함수나 명령을 동시에 실행할 때 사용

### 💡 컨텍스트 스위칭 비용 X

원래 CPU는 한 번에 하나의 스레드만 처리할 수 있어서, 효율성을 위해 여러 스레드를 전환하면 ‘컨텍스트 스위칭’ 비용이 발생한다. 적정 개수를 넘어 한 번에 너무 많은 스레드 수행하게 되면 성능이 저하되어 주의해야 한다.   

하지만 Go 언어에서는 CPU 코어마다 OS 스레드를 하나만 할당해서 사용하므로, 컨텍스트 스위칭 비용이 발생하지 않는다.

이를 가능하게 해주는 것이 고루틴이다.

- 스레드란?
- 고루틴/ 스레드의 차이점

## **24.2 고루틴 (goroutine)**

모든 프로그램은 고루틴을 최소 하나 가짐 (→ 메인 루틴)   

`go 함수_호출`  //메인 루틴 외에 고루틴 추가 생성

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

- 결과

```go
가 1 나 2 다 라 3 마 4 바 5 사
```

- `PrintHangul()`과 `PrintNumbers()` 함수는 서로 다른 고루틴에서 실행 → 동시에 실행
- 메인루틴에서 3초간 대기하는 이유? ( `time.Sleep(3 * time.Second)` )
    
    → 메인 함수가 종료되면 아무리 많은 고루틴이 있더라도 프로그램이 즉시 종료되므로, 기다려줘야 한다.     
    
    → 이를 해결할 수 있는 방법: sync 패키지의 WaitGroup 객체 사용
    

### 📌 서브 고루틴이 끝날 때까지 대기하기: `var wg sync.WaitGroup`

```go
var wg sync.WaitGroup

wg.Add(3)   //작업 개수 설정
wg.Done()   //작업 완료될 때마다 호출
wg.Wait()   //모든 작업이 완료될 때까지 대기
```

## **24.5 뮤텍스를 이용한 동시성 문제 해결**

뮤텍스 이용 → 고루틴에서 값을 변경할 때 다른 고루틴이 건들지 못하도록 통제

### 뮤텍스를 이용한 자원 접근 권한 통제 방식

- Lock() 메서드 호출 → 획득 or 대기
- Unlock() 메서드 호출 → 반납

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
	mutex.Lock()  //뮤텍스 획득
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
- DepositAndWithdraw() 함수에서 `mutex.Lock()` : **뮤텍스 획득**
- `defer mutex.Unlock()` : **반납**
- 잔고는 절대 0원 이하로 내려가지 않게 된다.
