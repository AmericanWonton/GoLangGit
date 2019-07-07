package main

import "fmt"

func main() {
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
}
