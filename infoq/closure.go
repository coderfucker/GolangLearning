package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Printf("sum: %v\n", sum(1, 2))

	sm := sumClosure()
	fmt.Println(sm())
	fmt.Println(sm())
	fmt.Println(sm())
	fmt.Println(sm())
}

func sumClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}