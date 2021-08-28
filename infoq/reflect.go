package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City string
}

type Person struct {
	Name string
	Age uint
	Address
}

func (p Person) Hello() {
	fmt.Println("Joe")
}

func main() {
	p := Person{"Joe", 24, Address{"Tokyo"}}

	t := reflect.TypeOf(p)
	fmt.Println("t:", t)

	fmt.Println("类型的名称：", t.Name())

	v := reflect.ValueOf(p)
	fmt.Println("v:", v)

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i).Interface()

		fmt.Printf("第%d个字段是：%s: %#v = %v \n", i + 1, key.Name, key.Type, value)
	}
}