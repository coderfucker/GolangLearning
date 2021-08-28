package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be greater than zero")
	}

	sum = a + b
	err = nil

	return
}

func sumMulti(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}

	return sum
}

func main() {
	res, ok := sum(1, 4)
	fmt.Printf("sum: %v ok: %v\n", res, ok)

	fmt.Printf("sumMulti: %v\n", sumMulti(1, 2, 3, 4))
}
