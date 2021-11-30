package main

import "fmt"

func main() {
	// 1) 큰따옴표로 묶기
	str1 := "Hello\t'World'\n"
	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를\n"
	// 2) 백쿼트로 묶기
	str2 := `Hello\t"Go"\nWorld!`
	poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println()
	fmt.Println(poet1)
	fmt.Println(poet2)

}
