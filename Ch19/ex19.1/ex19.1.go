package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { // 일반 함수
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // 메서드
	a.balance -= amount
}

func main() {
	a := &account{100}

	withdrawFunc(a, 30)

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)
}
