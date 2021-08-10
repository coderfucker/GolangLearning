package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFun(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(v)

	ScaleFun(&v, 10)
	fmt.Println(v)

	p := &Vertex{4, 3}
	p.Scale(3)
	fmt.Println(p)

	ScaleFun(p, 8)
	fmt.Printf("%v\n", p)

	fmt.Println(v, p)
}
