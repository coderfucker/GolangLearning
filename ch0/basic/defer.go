package main

import "fmt"

func main()  {
	defer fmt.Println("world")
	defer fmt.Println("我的世界")

	fmt.Println("hello")
}