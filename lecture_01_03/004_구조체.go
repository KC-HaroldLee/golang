package main

import(
	"fmt"
)

type hotdog int
type person struct{
	Fname string
	lname string // 소문자로 표기하면 패키지 밖에서 보이지 않는다고 한다(?), Private이라는 건지?
	Age int
}

func main() {
	var t hotdog // new같은 것인가
	t = 7
	fmt.Println(t)
	fmt.Printf("%T",t) // main.hotdog 라고 뜸 / {package명}.{타입명} < class는 아닐것 같다. 나중에 class는 어떻게 뜨는지 궁금하군
	fmt.Println("") // 개행

	p1 := person{
		"Harold",
		"Lee",
		35, // 아니..."," 찍지 않으면 에러다.
	}
	fmt.Println(p1)
	fmt.Println(p1.Fname) // 당연히 하나씩뽑는 것도 된다
	fmt.Println(p1.lname)
	fmt.Println(p1.Age)
	fmt.Printf("%T", p1)
	fmt.Println("") // 개행
}

