//JSON Marshalling
import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

/*
	The fields are uppercase in order to export them with Marshall
*/
p1 := person{
	First: "James",
	Last:  "Bond",
	Age:   32,
}
p2 := person{
	First: "Ugly",
	Last:  "Ass",
	Age:   45,
}

people := []person{p1, p2}
fmt.Println(people)

biteslice, err := json.Marshal(people)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(biteslice))

//JSON UNMARSHALLING
//Here's a good-ass site to see what Go code
//looks like for JSON objects: https://mholt.github.io/json-to-go/

import (
	"encoding/json"
	"fmt"
)

type person []struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

//Here's the JSON. You need those 'weird' accent marks, on the tilda key
var s string = `[{"First":"James","Last":"Bond","Age":32},{"First":"Ugly","Last":"Ass","Age":45}]`
biteSlice := []byte(s)
fmt.Printf("%T\n", s)
fmt.Printf("%T\n", biteSlice)

//people := []person{} This works too
var people []person

err := json.Unmarshal(biteSlice, &people)
if err != nil {
	fmt.Println(err)
}

fmt.Println("All of the data", people)
for i, v := range people {
	fmt.Println("\n PERSON NUMBER ", i)
	fmt.Println(v.First, v.Last, v.Age)
}