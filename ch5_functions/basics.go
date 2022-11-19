package main

import "fmt"

func main() {
	fmt.Println(div1(5, 2))
	fmt.Println(div1(6, 3))
	fmt.Println(div1(6, 0))

	fmt.Println(div2(5, 2))
	fmt.Println(div2(6, 3))
	fmt.Println(div2(6, 0))
}

func div1(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func div2(numerator, denominator float32) float32 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
