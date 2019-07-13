package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)        //This is the address of this variable
	fmt.Printf("%T\n", a)  //an int type
	fmt.Printf("%T\n", &a) //Pointer to an int type
}
