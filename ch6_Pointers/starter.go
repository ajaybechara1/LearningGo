package main

import "fmt"

func main() {
	str := "ajay"
	pointerStr := &str

	fmt.Println(pointerStr)

	str2 := "bechara"
	pointerStr2 := &str2

	fmt.Println(pointerStr2)

	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // prints nil
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}
