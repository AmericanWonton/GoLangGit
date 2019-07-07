package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

/*This is a function method. Any struct/type can use this if it is
put in the parameter field after 'func'. Think of it like class
methods.
*/
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Uglyass",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
