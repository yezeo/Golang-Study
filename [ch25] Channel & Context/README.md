# **Chapter 25. Channel & Context**

## **25.1 채널 사용하기**

**채널**이란? 고루틴끼리 메시지를 전달할 수 있는 메시지 큐

### 📌 **채널 인스턴스 생성**

- 채널 키워드 : `chan`
- 생성 방법: `make(chan 자료형)`

```go
var messages chan string = make(chan string)
```

### 📌 **채널에 데이터 넣기:** `변수 ←`

- 좌변: 채널 인스턴스 / 우변: 넣을 데이터
- `messages <- "This is message"`

### 📌 **채널에서 데이터 빼기:**  `← 변수`

- 좌변: 빼낸 데이터를 담을 변수 / 우변: 채널 인스턴스
- `var msg string = <- messages`

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup  //wg 변수 생성
	ch := make(chan int) //채널 생성

	wg.Add(1)
	go square(&wg, ch) //고루틴 생성
	ch <- 9   //채널에 데이터 넣기
	wg.Wait() //작업이 완료되길 기다림
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch  // 채널에서 데이터 빼기

	time.Sleep(time.Second) //1초 대기
	fmt.Printf("Square: %d\n", n*n)  // 9*9
	wg.Done()
}
```

- 결과: `Square: 81`

### 📌 **채널 크기**

- 일반적으로 채널 생성 시 크기= 0
    
    → 데이터를 넣을 때 보관할 곳이 없어서, 데이터를 빼갈 때까지 대기 (그 전엔 프로그램 실행 X)
    

```go
package main

import "fmt"

func main() {
	ch := make(chan int) //크기 0인 채널

	ch <- 9
	fmt.Println("Never print") //실행되지 않는다
}
```

- 결과

`fatal error: all goroutines are asleep - deadlock!`

### 📌 **버퍼를 가진 채널**

- **버퍼**: 내부에 데이터를 보관할 수 있는 메모리 영역
- **버퍼를 가진 채널** = **보관함을 가진 채널**

```go
var chan string messages = make(chan string, 2)   //버퍼 2개인 채널 생성
```

### 📌 **채널에서 데이터 대기**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
```

- 결과

```go
Square: 0
Square: 4
Square: 16
Square: 36
Square: 64
Square: 100
Square: 144
Square: 196
Square: 256
Square: 324
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
...
```

- `wg.Wait()` 메서드로 작업 완료될 때까지 기다림 → for range 구문은 채널에 데이터가 들어오기를 기다림 → 모든 고루틴이 멈춤 (**deadlock**)
- 채널을 다 사용하면 close(ch) 호출 → 채널이 모두 빈 상태에서 닫혔으면 for range문을 빠져나감

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}
```

- 결과

```go
Square: 0
Square: 4
Square: 16
Square: 36
Square: 64
Square: 100
Square: 144
Square: 196
Square: 256
Square: 324
```

### 📌 **select문**

select문은 언제 쓸까?   

- 채널에서 데이터를 대기하는 상황에서, 만약 데이터가 들어오지 않으면 다른 작업을 하고 싶을 때
- 여러 채널을 동시에 대기하고 싶을 때

```go
select{
case n := <-ch1:
...  //ch1 채널에서 데이터를 빼낼 수 있을 때 실행
case n := <-ch2:
...  //ch2 채널에서 데이터를 빼낼 수 있을 때 실행
case ...
}
```

예제

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}
```

- 결과

```go
Square: 0
Square: 4
Square: 16
Square: 36
Square: 64
Square: 100
Square: 144
Square: 196
Square: 256
Square: 324
```

- quit 종료 채널 : square() 루틴을 만들 때 알려줌
- select문 : ch와 quit 채널 모두 기다림 (ch 채널에서 데이터를 읽을 수 있다면 계속 읽음)

### 📌 **일정 간격으로 실행**

메시지 있다면 빼와서 실행, 없다면 1초 간격으로 다른 일을 수행한다고 가정

time 패키지의 `Tick()` 함수로 원하는 시간 간격마다 신호를 보내주는 채널 생성 가능

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
```

- 결과

```go
Square: 0
Square: 4
Tick
Square: 16
Square: 36
Tick
Square: 64
Square: 100
Square: 144
Tick
Square: 196
Square: 256
Square: 324
Tick
Terminated!
```

- select문을 이용해서 tick, terminate, ch 순서로 채널에서 데이터 읽기 시도
- tick은 1초 간격으로 신호 전송, 10초 이후에는 terminate 신호 (함수 종료)

## **25.2 컨텍스트 사용하기**

컨텍스트(Context)란? context 패키지에서 제공하는 기능.   

작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할

### 📌 **작업 취소가 가능한 컨텍스트**

지시자가 원할 때 작업 취소를 알릴 수 있다.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
```

- 결과

```go
Tick
Tick
Tick
Tick
Tick
```

### 📌 **작업 시간을 설정한 컨텍스트**

일정한 시간 동안만 작업을 지시할 수 있는 컨텍스트   

`ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)`   

context 패키지의 `WithTimeOut()` 함수 : 작업 시간 설정 (시간이 지난 뒤 컨텍스트의 Done() 채널에 시그널 전송 = 작업 종료 요청)

### 📌 **특정 값을 설정한 컨텍스트**

작업자에게 작업을 지시할 때 별도 지시사항을 추가하고 싶을 때 사용

```go
package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
```

- 결과 `Square:81`
- “number”를 키로 값을 9로 설정한 컨텍스트 생성
- 컨텍스트의 `Value()` 메서드로 값을 읽어옴

컨텍스트에 값을 설정해서 다른 고루틴으로 작업을 지시할 때 외부 지시사항으로 설정 가능, 지시자와 작업자 사이에 어떤 키로 어떤값이 들어올지 약속
