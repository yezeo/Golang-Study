package main

import "fmt"

func main() {
	a := 3.5
	var b int = int(a * 3) // b= int(10.5) = 10
	var c int = int(a) * 3 // c= 3*3 = 9
	var isEqual bool = (b == c) // 두 값이 같은가

	var typeBig int16 = 3456
	var typeSmall int8 = int8(typeBig)

	fmt.Println(isEqual) // false
	fmt.Println(b, c)  // 10 9

	fmt.Println(typeBig)
	fmt.Println(typeSmall)
}
