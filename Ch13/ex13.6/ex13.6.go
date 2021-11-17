package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
}
