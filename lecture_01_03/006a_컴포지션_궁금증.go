package main

import(
	"fmt"
)

type inner struct{ 
	name string
}

type outer struct{
	inner
	name string 
}

func (i inner) speak() {
	fmt.Println(i.name, `is inner name`) 
}
func (o outer) speak() {
	fmt.Println(o.name, `is outer name`) 
}


func main() {
	inner_1 := inner{
					"Harold",
				}

	outer_1 := outer {
		inner_1,
		"Lee",
	}
	fmt.Println(inner_1) //{Harold}
	fmt.Println(outer_1) //{{Harold} Lee}
	fmt.Println(inner_1.name) //Harold
	fmt.Println(outer_1.name) //Lee
}	
