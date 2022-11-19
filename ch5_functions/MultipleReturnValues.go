package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("result", result, " remainder", remainder)

	result, remainder, err = divAndRemainder(5, 2)
	fmt.Println("result", result, " remainder", remainder)
	result, remainder, err = divAndRemainder2(5, 2)
	fmt.Println("result", result, " remainder", remainder)
	result, remainder, err = divAndRemainder3(5, 2)
	fmt.Println("result", result, " remainder", remainder)
	result, remainder, err = divAndRemainder4(5, 2)
	fmt.Println("result", result, " remainder", remainder)
}

func divAndRemainder2(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder3(numerator int, denominator int) (a1 int, a2 int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	result, remainder := numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder4(numerator, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}
