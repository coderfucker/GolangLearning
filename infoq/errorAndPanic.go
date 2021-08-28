package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("10a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
