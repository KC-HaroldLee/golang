package main

import(
	"fmt"
)

type person struct{ // 소문자로 돌려 놓음
	fname string
	lname string
	age int
}

type secretAgent struct{
	person // 당연히 된다
	licenseToKill bool
}

// 다시 봐도 참 이상한 구조
// func (receiver) identifier(parameters) (returns) {<code>}
func (p person) speak() {
	fmt.Println(p.fname, `says, Good morning!`)
}
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "No way!"`)
}


func main() {
	person_1 := person{
					"Harold",
					"Lee",
					35, }
	sa_1 := secretAgent {
		person_1,
		true,
	}

	// 각 구조체의 메서드를 만드는 느낌에 가깝다.
	person_1.speak()
	sa_1.speak() // 재밌는건 하위 구조체를 언급하지 않아도 되었다는 점
	sa_1.person.speak() // 이것도 가능
}	