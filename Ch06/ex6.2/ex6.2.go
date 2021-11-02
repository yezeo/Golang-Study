package main

import "fmt"

func main() {
	var x1 int8 = 34   //8비트 정수, 0010 0010
	var x2 int16 = 34  //16비트 정수, 0000 0000 0010 0010
	var x3 uint8 = 34  //8비트 부호 없는 정수
	var x4 uint16 = 34 //16비트 부호 없는 정수

	fmt.Printf("^%d = %5d, \t %08b\n", x1, ^x1, uint8(^x1)) //➎
	fmt.Printf("^%d = %5d, \t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d, \t %016b\n", x4, ^x4, ^x4)

}
