package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var myHouse House
	myHouse.Address = "Gwangmyeong"
	myHouse.Size = 32
	myHouse.Price = 9.8
	myHouse.Type = "Apartment"

	fmt.Println("주소: ", myHouse.Address)
	fmt.Printf("크기: %d평", myHouse.Size)
	fmt.Printf("가격: %.2f억 원\n", myHouse.Price)
	fmt.Println("타입: ", myHouse.Type)
}
