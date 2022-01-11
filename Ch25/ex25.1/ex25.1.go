package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) //채널 생성

	wg.Add(1)
	go square(&wg, ch) //Go 루틴 생성
	ch <- 9
	wg.Wait() //작업이 완료되길 기다린다
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second) //1초 대기
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
