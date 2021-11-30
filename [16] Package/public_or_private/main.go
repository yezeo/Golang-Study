package main

import (
	"ch16/public_or_private/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
}
