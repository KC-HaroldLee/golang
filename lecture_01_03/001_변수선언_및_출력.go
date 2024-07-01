// go 설치 참고하기 좋은 곳
// https://snowdeer.github.io/go/2018/01/19/how-to-install-golang-on-ubuntu/

package main //언젠가 이 비밀도 벗겨지겠지

import(
	"fmt" // 아 맞다 따옴표
)

func main() {
	x := 7
	// fmt.Println(x) // 단순 개행 프린트
	fmt.Printf("%T", x) // 개행이 없다.
}

