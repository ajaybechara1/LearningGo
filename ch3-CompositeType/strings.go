package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var s string = "Hello, 🌞"
	fmt.Println("len of s: ", len(s))
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs, len(bs), binary.Size(rs))
	fmt.Println(rs, len(rs), binary.Size(rs))
}
