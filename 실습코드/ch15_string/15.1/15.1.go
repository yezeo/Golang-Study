package main

import "fmt"

func main() {
	str1 := "Hello\t'World'\n"

	str2 := `Go is "awesome"\n Go is simple and\t 'powerful'`

	fmt.Println(str1)
	fmt.Println(str2)
}
