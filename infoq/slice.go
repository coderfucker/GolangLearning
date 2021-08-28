package main

import "fmt"

func main() {
	slice := make([]string, 4, 8)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	slice[3] = "d"
	slice1 := append(slice, "e")

	slice2 := []string{"a", "b", "c", "d"}
	slice3 := append(slice2, "e")
	slice4 := append(slice3, "x", "y", "z", "o")

	fmt.Printf("slcie: %v\n", slice)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice2, slice3, slice4, cap(slice2), cap(slice3), cap(slice4))
}
