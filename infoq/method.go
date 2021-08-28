package main

import "fmt"

type Name string

func (n Name) String() {
	fmt.Println("name is ", n)
}

func main() {
	name := Name("hello")
	name.String()
}