package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("haha")
	defer fmt.Println("hehe")

	fmt.Println("hello")
}