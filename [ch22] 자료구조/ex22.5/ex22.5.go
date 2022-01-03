package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["김윤서"] = "경기도 광명시"
	m["이선주"] = "대전광역시 봉명동"
	m["윤창용"] = "서울특별시 염창동"

	m["윤창용"] = "서울특별시 상암동"

	fmt.Printf("김윤서의 주소는 %s입니다.", m["김윤서"])
	fmt.Printf("윤창용의 주소는 %s입니다.", m["윤창용"])
}
