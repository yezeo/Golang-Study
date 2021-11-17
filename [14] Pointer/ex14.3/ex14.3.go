package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// 단순 값 복사 -> 실제 main에서 변경되지 않는 문제
func ChangeData(d Data) {
	d.value = 200
	d.data[100] = 999
}

// 포인터 사용 -> 값 변경 가능
func ChangeData_withPointer(d *Data) {
	d.value = 200
	d.data[100] = 1000
}
func main() {
	var data Data
	data.value = 0
	data.data[100] = 100

	ChangeData(data) // 인수로 data를 넣는다.

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100]) // data 필드 출력

	ChangeData_withPointer(&data) //인수로 data의 주솟값을 넣는다.

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100]) // data 필드 출력
}
