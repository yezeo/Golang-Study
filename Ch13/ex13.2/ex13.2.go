package main

import "fmt"

type User struct { // 일반 고객용 구조체
	Name string
	ID   string
	Age  int
}

type VIPUser struct { // VIP 고객용 구조체
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwrang", 40},
		3,
		25,
	} // User를 포함한 VIPUser 구조체 변수 초기화

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d VIP 가격: %d만원\n",
		vip.UserInfo.Name, // UserInfo 안의 Name
		vip.UserInfo.ID,   // UserInfo 안의 ID
		vip.UserInfo.Age,
		vip.VIPLevel, // IPUser의 VIPLevel
		vip.Price,
	)
}
