package main

import "fmt"

func fabonacci() func() int {
	var s []int
	return func() int {
		if len(s) < 2 {
			s = append(s, len(s))
		} else {
			s = append(s, s[len(s)-1] + s[len(s)-2])
		}
		return s[len(s) - 1]
	}
}
func main() {
	f := fabonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
