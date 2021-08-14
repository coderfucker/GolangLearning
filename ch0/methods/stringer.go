package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string  {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {
	a := Person{"Lee", 42}
	b := Person{"James", 32}

	fmt.Println(a, b)
}
