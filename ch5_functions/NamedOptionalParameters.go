package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	return nil
}

func main() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

	// In practice, not having named and optional parameters isn’t a limitation.
	// A function shouldn’t have more than a few parameters, and named and optional parameters are mostly useful
	// when a function has many inputs. If you find yourself in that situation,
	// your function is quite possibly too complicated.
}
