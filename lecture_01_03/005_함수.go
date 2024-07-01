package main

import(
	"fmt"
)

type person struct{
	Fname string
	lname string // 소문자로 표기하면 패키지 밖에서 보이지 않는다고 한다(?), Private이라는 건지?
	Age int
}

// 참 이상한 구조
// func (receiver) identifier(parameters) (returns) {<code>}

func (p person) speak() {
	fmt.Println(p.Fname, `says, Good morning!`)
}

func main() {
	person_1 := person{
					"Harold",
					"Lee",
					35, }
	// 근데 함수 호출 방식이 특이하다.
	// 파라미터를 넣어봐야 확실히 알듯
	person_1.speak()
	// fmt.Println(person_1)
}	