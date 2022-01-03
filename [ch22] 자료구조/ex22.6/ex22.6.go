package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product) //key가 int, value가 구조체인 맵

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"필통", 1000}
	m[345] = Product{"샤프", 700}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
