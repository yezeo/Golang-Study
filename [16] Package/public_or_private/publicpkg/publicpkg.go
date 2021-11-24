package publicpkg

import "fmt"

const (
	PI = 3.1415   // 공개
	pi = 3.141516 //공개 x
)

var ScreenSize int = 1080
var screenHeight int

func PublicFunc() {
	const MyConst = 100
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() {
	fmt.Println("This is a private function")
}

type MyInt int
type myString string

type MyStruct struct {
	Age  int
	name string
}

func (m MyStruct) PublicMethod() {
	fmt.Println("This is a public method")
}

func (m MyStruct) privateMethod() {
	fmt.Println("This is a private method")
}

type myPrivateStruct struct {
	Age  int
	name string
}

func (m myPrivateStruct) PrivateMethod() {
	fmt.Println("This is a private method")
}
