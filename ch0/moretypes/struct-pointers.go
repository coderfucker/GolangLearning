package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	(*p).Y = int(math.Pow(2, 10))
	fmt.Println(v)
}