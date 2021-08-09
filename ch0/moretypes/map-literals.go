package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex {
		40.6833, -74.39967,
	},
	"Google": {
		Lat: 37.42202, Long: -122.08408,
	},
}

func main() {
	fmt.Println(m, m["Google"])
}