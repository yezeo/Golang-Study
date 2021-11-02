package main

import "fmt"

func main() {
	var x int8 = 16   //부호 비트값이 0인 수
	var y int8 = -128 //부호 비트값이 1인 수
	var z int8 = -1   //모든 비트값이 1인 정수
	var w uint8 = 128 //최상위 비트값이 1인 양수
	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", uint8(y), uint8(y>>2), y>>2)
	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), w>>2)
}
