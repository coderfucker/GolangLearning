package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}