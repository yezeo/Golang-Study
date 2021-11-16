package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}
type VIP struct {
	UserInfo User // User 구조체를 포함
	Level    int
	Price    int
}

func main() {
	user := User{"Yoonseo", "dotsi", 22}
	vip := VIP{
		User{"Eunxung", "ae", 22},
		3,
		250,
	}
	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d VIP 가격: %d만원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.Level,
		vip.Price,
	)
}
