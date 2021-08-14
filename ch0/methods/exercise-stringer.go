package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var arrSlice []string
	for _, value := range ip {
		arrSlice = append(arrSlice, fmt.Sprint(value))
	}

	return fmt.Sprintf("%v", strings.Join(arrSlice, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}