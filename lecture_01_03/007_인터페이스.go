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

type human interface{
	speak()
}

// 이것도 가능하군
func saySomething(h human){
	h.speak()
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
	saySomething(person_1) // Harold says, Good morning!
	saySomething(sa_1) // Harold Lee says, "No way!"
}	
