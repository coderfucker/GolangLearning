package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v * 2)
	case string:
		fmt.Printf("%q is %v bypes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var m = make(map[string]int)
	m["age"] = 18

	var arr = [2]string{"a", "b"}

	do(21)
	do("Hello")
	do(m)
	do(arr)
}