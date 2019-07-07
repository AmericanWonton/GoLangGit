//Funcitons
testFunc()

//func (r reciever) identifier(parameters) (return(s))

func testFunc() {
fmt.Println("Hello from testFunc")
}

testFunc2("Kill yourself")

func testFunc2(aMessage string) {
fmt.Printf("Here's your message: %v\n", aMessage)
}

s1 := returnTestFunc("Big peener")
fmt.Println(s1)

func returnTestFunc(takeString string) string {
return fmt.Sprint("Hello from returnTestFunc,", takeString)
}

x, y := anotherTest("Big", "Oofus") //returns a string to x and bool to y
fmt.Println(x)
fmt.Println(y)

func anotherTest(fn string, ln string) (string, bool) {
a := fmt.Sprint(fn, ln, `, says "Hello" `)
b := false
return a, b
}

//Varadic Parameter
x := sum(1, 2, 3, 4, 5, 6, 7)
fmt.Println("The total is ", x)

func sum(x ...int) int { //Infinite parameters
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("For item in position %v, we are now adding", i)
		fmt.Printf("%v to the total, which is now %v\n", v, sum)
	}
	return sum
}

//"Unfurling" a slice
xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
x := sum(xi...) //These are the Ints "unfurled" into the parameter
//you could also pass in 0 arguments x := sum()
fmt.Println("The total is ", x)

func sum(x ...int) int { //Infinite parameters
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("For item in position %v, we are now adding", i)
		fmt.Printf("%v to the total, which is now %v\n", v, sum)
	}
	return sum
}

//Defer
defer foo() //Runs at the END of the func main
//Good for closing file
bar()

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}

//Method