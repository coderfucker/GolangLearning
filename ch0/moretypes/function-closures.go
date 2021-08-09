package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("sum:", sum)

	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, _ := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}