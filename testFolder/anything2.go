package testFolder

testVal := "It's a test value"
testInt := 22
testFloat := 55.63
testBool := false

//bit shifting

//iota and bit shifting
const (
	a = iota
	b = iota
	c = iota
)
//bit shifting 
const (
	_ = iota
	//kb = 1024
	kb = 1 << (iota * 10)
	//mb = kb * 1024
	mb = 1 << (iota * 10)
	//gb = 1024 * mb
	gb = 1 << (iota * 10)
)
fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
//Type
type bigPenis string
type biggerPenis bigPenis

var x bigPenis

//Array
var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x)) //The length of array x

//composite literal
x := []int{4, 5, 7, 8, 42}
fmt.Println(x)

//Slice allows you to group values of the same type
x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(" ")
	fmt.Println(" ")

	for i := 0; i < (len(x)); i++ {
		fmt.Println(x[i])
	}

	//slicing a slice
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println("Beneath you is the sliced slice, dumbass")
	fmt.Println(x[1:])
	fmt.Println("Here's another slice of x..up to, but not including...")
	fmt.Println(x[1:3])
	//appending a slice to a slice
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println("Yo, we about to append a slice to slice 'x'")
	x = append(x, 77, 88, 200, 666)
	fmt.Println(x)

	y := []int{234, 456, 678, 999}
	x = append(x, y...) //The dots mean you append at the end.
	fmt.Println(x)

	//deleting from slice
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 200, 666)
	fmt.Println(x)

	y := []int{234, 456, 678, 999}
	x = append(x, y...) //The dots mean you append at the end.
	fmt.Println(x)

	/*delete from start of one place, resume at other place to end
	 */
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//Make Slice
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x)) //current length
	fmt.Println(cap(x)) //capacity of slice
	x[0] = 42
	x[9] = 999

	x = append(x, 3434)
	fmt.Println(len(x)) //current length
	fmt.Println(cap(x)) //capacity of slice
	fmt.Println(x)

	//Multidemensional slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "strawberry", "hazelnut"}
	fmt.Println(mp)

	fmt.Println("  ")

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	//Maps
	uglyass := "uglyass"
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
		uglyass:           88, //DON'T FORGET THE COMMAS
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Barnabas"]) //This shit aint here, it prints 0
	//Here's how to check if it exists
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
	//Here's the check in an 'if' statement
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("Here's your lost value", v)
	}
	if v, ok := m[uglyass]; ok {
		fmt.Println("Here's uglyass", v)
	}

	//Maps, add element in range
	m := map[string]float64{
		"Cocaine": 43.50,
		"Weed":    20.50,
		"Heroin":  8.75,
		"Crack":   45.56,
	}
	fmt.Println(m)

	m["Super Cocaine"] = 70.8

	for k, v := range m {
		fmt.Println(k, v)
	}
	//delete map entry
	m := map[string]float64{
		"Cocaine": 43.50,
		"Weed":    20.50,
		"Heroin":  8.75,
		"Crack":   45.56,
	}
	fmt.Println(m)

	delete(m, "Crack")
	fmt.Println(m)

	//Structs
	type person struct{
		first string
		last string
	}

	p1 := person{
		first: "James",
		last: "Bond",
	}
	
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p2.first)
	fmt.Println(p2.last)

	//Embedded Struct
	type person struct {
		first string
		last  string
		age   int
	}
	type secretAgent struct {
		person       //This DOSEN'T NEED A FIELD NAME, JUST THE TYPE
		license2Kill bool
	}

	sam := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		license2Kill: true,
	}

	fmt.Println(sam)
	fmt.Println(sam.person.age)

	//Anonymous Structs
	person1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Ugly",
		last:  "ass",
		age:   42,
	}

	fmt.Println(person1)
