package main

import "fmt"

type House struct { // House 구조체 정의
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House               // House 구조체 변수 선언
	house.Address = "서울시 강동구 ..." // 각 필드값 초기화
	house.Size = 28
	house.Price = 98
	house.Type = "아파트"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기 %d평\n", house.Size)
	fmt.Printf("가격: %.2f억원\n", house.Price)
	fmt.Println("타입:", house.Type)
}
