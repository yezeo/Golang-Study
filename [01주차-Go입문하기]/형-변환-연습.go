package main

import "fmt"

func main() {
  a:=3 // var a int = 3 과 동일
  var b float64 = 3.5 //float64형
  
  var c int = int(b)  // float->int 변환 시 소수점 버리고 3이 됨
  d := float64(a*c)
  
  var e int64 = 7
  f := int64(d) * e
  
  var g int = int(b*3)
  var h int = int(b) * 3
  fmt.println(g, h, f)  // 10 9 63
}
