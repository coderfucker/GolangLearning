package main

import (
	"math"
	"fmt"
)

func main() {
	_, v, _ := getData()
	print(v)
	println(v)
	fmt.Println(v)
	
	f := pow(2, 10, 512)
	println(f)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	total := 1
	for total < 100 {
		total += total
	}

	fmt.Println(total)

	for {
		println("smida.che")
	}
}

func getData() (int, int, int) {
	return 2, 4, 8
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}