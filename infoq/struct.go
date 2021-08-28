package main

import "fmt"

type address struct {
	city string
}

type person struct {
	name string
	age uint
	addr address
}

type Info interface {
	GetInfo() string
}

func (p person) GetInfo() string {
	return fmt.Sprintf("my name is %s, age is %d, addr is %s", p.name, p.age, p.addr.city)
}
//func (p *person) GetInfo() string {
//	return fmt.Sprintf("my name is %s, age is %d, addr is %s", p.name, p.age, p.addr.city)
//}
func printInfo(i Info) {
	fmt.Println(i.GetInfo())
}

func main() {
	p := person{
		name: "Joe",
		age: 24,
		addr: address{"Tokyo"},
	}

	fmt.Println(p, p.name, p.addr.city)

	//printInfo(p)
	printInfo(&p)
}
