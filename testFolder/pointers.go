//Pointers
a := 42
	fmt.Println(a)
	fmt.Println(&a)        //This is the address of this variable
	fmt.Printf("%T\n", a)  //an int type
	fmt.Printf("%T\n", &a) //Pointer to an int type

	var b *int = &a //variable 'b' is a pointer to an int,(&a)
	fmt.Println(b)
	fmt.Println(*b)  //The asteriks is part of the 'type'
	fmt.Println(*&a) //asteriks give you value of the address

	*b = 43 //since value 'b' has the address of a, it changes the
	//value of 'a'
	fmt.Println(a)
/*
Pointers are good for parsing big values and changing values at an
address.
*/

x := 0
testFunction(&x)
fmt.Println(x)

func testFunction(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 93
	fmt.Println(y)
	fmt.Println(*y)
}

//Method Sets
