package main

import "fmt"

func main() {
	// Complete for Statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Condition only for Statement
	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 2
	}

	// The infinite for Statement
	for {
		fmt.Println("Hello! this is infinite for loop 😀")
		break
	}

	// break and continue is as normal

	// for-range loop
	// with slice
	evenValues := []int{2, 4, 6, 8, 10, 12, 14}
	for i, v := range evenValues {
		fmt.Println(i, v)
	}

	for _, v := range evenValues {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// iterating over maps
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// iterating over string
	samples := []string{"hello 🙋a🙋", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// First, notice that the first column skips the number 7. Second, the value at position 6 is 960.

	// What we are seeing is special behavior from iterating over a string with a for-range loop.
	// It iterates over the runes, not the bytes. Whenever a for-range loop encounters a multibyte rune in a string,
	// it converts the UTF-8 representation into a single 32-bit number and assigns it to the value.
	// The offset is incremented by the number of bytes in the rune.
	// If the for-range loop encounters a byte that doesn’t represent a valid UTF-8 value,
	// the Unicode replacement character (hex value 0xfffd) is returned instead.

	// Use a for-range loop to access the runes in a string in order.
	// The key is the number of bytes from the beginning of the string, but the type of the value is rune.

	samples = []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		fmt.Println("start of outer loop")
		for i, r := range sample {
			// process innerVal
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		// here we have code that runs only when all the
		// innerVal values were successfully processed
		fmt.Println("after inner for loop")
	}

	test := "Hello 😀"
	for i := 0; i < len(test); i++ {
		fmt.Println(string(test[i]))
	}
}
