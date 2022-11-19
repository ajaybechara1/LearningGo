package main

type opFuncType func(int, int) int

var opMap2 = map[string]opFuncType{}

func main() {
	opMap2["+"] = func(i int, i2 int) int {
		return i + i2
	}
}
