package main

import(
	"fmt"
)

func main() {
	xi := []int{2,4,6,8,10} // 자바와 비슷
	fmt.Println(xi)

	// map도 가능
	// 1:1 매칭임을 기억하고 작성방법을 떠올리자
	mi := map[string]int{
		"석현" : 123,
		"인환" : 456,
		"진철" : 789, // 마지막 콤마를 허용한다
	}
	fmt.Println(mi) // 입력하지 않으면 "declared and not used"이라며 go run이 되지 않는다. 나중에 방법을 알게 되겠지
}

