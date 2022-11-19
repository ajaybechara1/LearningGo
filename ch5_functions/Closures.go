package main

import (
	"fmt"
	"sort"
)

// functions declared inside a function is called closure. This is computer science word that means functions declared
// inside functions area able to access and modify variables declared in th outer function

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	twoMult := makeMult(2)
	threeMult := makeMult(3)
	fmt.Println(twoMult(2))
	fmt.Println(threeMult(2))
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
