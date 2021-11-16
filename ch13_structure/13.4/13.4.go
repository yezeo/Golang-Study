package main

import "fmt"

type User struct {
	Name string
	ID string
	Age int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser {
		User{"화랑", "hwarang", 40,10}, 250, 3}
	}

	fmt.Printf("유저: %s ID: %s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이:%d VIP 레벨 : %d VIP 가격: %d만원 \n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)

}
