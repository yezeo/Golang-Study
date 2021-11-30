package main

import (
	"fmt"
	"strings"
)

func ToUpper(str string) string {
	var builder strings.Builder
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			builder.WriteRune('A' + (v - 'a'))
		} else {
			builder.WriteRune(v)
		}
	}
	return builder.String()
}

func main() {
	str := "hello World!"

	fmt.Println(ToUpper(str))
}
