package main

import "fmt"

func main() {
	//x := make([]int, 5, 20)
	////fmt.Println(cap(x))
	//x = append(x, 10)
	////fmt.Println(cap(x))
	//x = append(x, 7, 8, 9, 10, 11)
	////fmt.Println(cap(x))
	//
	//x = []int{1, 2, 3, 4, 5}
	//x = append(x, 6)
	////fmt.Println(x, cap(x))
	//y := x[:2]
	////fmt.Println(y, cap(y))
	//y = append(y, 30)
	////fmt.Println(y, cap(y))
	////fmt.Println(x, cap(x))
	//z := x[2:]
	////fmt.Println(z, cap(z))
	////z = append(z, 10, 11, 12, 13, 14, 15, 16, 17)
	////fmt.Println(x, cap(x))
	////fmt.Println(z, cap(z))
	//z[2] = 13
	////fmt.Println(x, cap(x))
	////fmt.Println(z, cap(z))

	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	//y := x[:2]
	//z := x[2:]
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	y = append(y, 30, 40, 50)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	x = append(x, 60)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	z = append(z, 70)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
