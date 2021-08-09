package main

import "fmt"

type Vertex struct {
	Lat, Long float64
	name string
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967, "che"}
	fmt.Println(m["Bell Labs"])
}