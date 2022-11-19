package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var ajay person
	fmt.Println(ajay)

	var jaydip = person{
		"Jaydeep",
		20,
		"dog",
	}
	fmt.Println(jaydip)

	var sanket = person{
		name: "Sanket",
		age:  20,
	}
	fmt.Println(sanket)

	// Anonymous Structs
	pet := struct {
		name string
		kind string
	}{name: "Fido", kind: "Dog"}
	fmt.Println(pet)
}
