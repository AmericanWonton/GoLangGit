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

//Interfaces & Polmorphism
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "the secret agent speak.")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " the person speak.")
}

type human interface {
	speak() //If you have the methods in here, anything that uses it
	//becomes the "type" of human
}

func bar(h human) {
	/*Below is an 'assertion' example...as in, we
	ASSERT, that this is totally the type specified, so
	access the values of that type.
	*/
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar, ", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar, ", h.(secretAgent).first)
	default:
		fmt.Println("I was passed into bar, ", h)
	}

}


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
p1 := person{
	first: "Dr",
	last:  "Yes",
}
fmt.Println(sa1)
sa1.speak()
sa2.speak()

fmt.Println(p1)

bar(sa1)
bar(sa2)
bar(p1)

//Anonymous func, no arguments
func() {
	fmt.Println("Anonymous func ran")
}()
//Anonymous func, with arguments
func(x int) {
	fmt.Println("Here's a number: ", x)
}(42)

//Func Expression
f := func() {
	fmt.Println("My first func expression")
}

f()

//Return function from a function
s := stringTest()

	fmt.Println(s)

	x := intTest() //Here, x is returning a function of int,
	//from intTest func
	fmt.Printf("%T\n", x)

	i := x() //This actually assigns the value of func int
	//,(from the above 'x' assignment) to the return value of 451.
	fmt.Println(i)

	//You could also do below...
	fmt.Println(intTest()())

	func stringTest() string {
		s := "Hello world"
		return s
	}
	func intTest() func() int {
		return func() int {
			return 451
		}
	}

	//Callbacks...passing in functions as parameters
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println("Even Numbers,", s2)

	func sum(xi ...int) int {
		//fmt.Printf("%T\n", xi)
		total := 0
		for _, v := range xi {
			total += v
		}
		return total
	}
	
	func even(f func(xi ...int) int, vi ...int) int {
		var yi []int
		for _, v := range vi {
			if v%2 == 0 {
				yi = append(yi, v)
			}
		}
		return f(yi...)
	}

	/*Closure-Limiting the scope to the variable*/
	 
	//Recursion-Maybe use a loop instead, but whatever
	//It's when a function calls itself...don't do it.

