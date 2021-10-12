package main

import "fmt"

var g int = 10 //global variable

func main() {
	var m int = 20 //local variable
	{
		var s int = 50 //local variable in {}
		fmt.Println(m, s, g)
	
	} // variable 's' dead

	fmt.Println(m, g)  // 20 10
	fmt.Println(s)  // ERROR
}
