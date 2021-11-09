// 음식값이 오만 원이 넘고 친구 중에 부자가 있다면 신발끈을 묶는다.
// 부자가 없다면 돈을 나눠 낸다.
// 음식값이 3만 원 이상, 오만 원 이하이고 같이 간 친구 수가 3명 초과이면 신발끈을 묶는다.
// 3명 이하이면 돈을 나눠 낸다.
// 3만 원 미만이면 내가 낸다.

package main

import "fmt"

// 친구 중 부자가 있는지 반환
func HasRichFriend() bool {
	return true
}

// 같이 간 친구 숫자를 반환
func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈을 묶는다")
		} else {
			fmt.Println("나눠 낸다")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("신발끈을 묶는다")
		} else {
			fmt.Println("나눠 낸다")
		}
	} else {
		fmt.Println("내가 낸다")
	}
}
