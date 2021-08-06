package main

import "fmt"

func main()  {
	a := make([]int, 5)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := make([]int, 0, 5)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = b[:cap(b)]
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = b[1:]
	fmt.Println(len(b))
	fmt.Println(cap(b))
}
