package main

import "fmt"

func main() {
	mapMake := make(map[string]int)
	mapLiteral := map[string]int{"a": 1}

	mapMake["a"] = 1
	mapMake["b"] = 10

	age, ok := mapMake["z"]

	for s, ok := range mapMake {
		fmt.Println(s)
		fmt.Println(ok)
	}

	s := "你好world"
	bs := []byte(s)
	for i, b := range bs {
		fmt.Printf("bs unicode: %v - %v\n", i, b)
	}

	fmt.Println(mapMake, age, ok, len(mapMake))
	fmt.Println(mapLiteral["a"])
}
