package main

const Y int = 3 //상수

func main() {
	var x int = 5              //변수
	a := [x]int{1, 2, 3, 4, 5} // ERROR >> 변수를 배열크기로 사용할 수 없음

	b := [Y]int{1, 2, 3} // OK

	var c [10]int // OK

	//참고. 변수를 선언하고 사용하지 않으면 컴파일 에러가 난다!!!
}
