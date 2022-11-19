package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	printXyz(x, y, z)
	//y = append(y, 30, 40, 50)
	//x = append(x, 60)
	//z = append(z, 70)
	//printXyz(x, y, z)
	x[2] = 1000
	printXyz(x, y, z)
}

func printXyz(x []int, y []int, z []int) {
	fmt.Printf("x: %p ", &x)
	fmt.Println(x)
	fmt.Printf("y: %p ", &y)
	fmt.Println(y)
	fmt.Printf("z: %p ", &z)
	fmt.Println(z)
	fmt.Println("cap: ", cap(x), cap(y), cap(z))
}
