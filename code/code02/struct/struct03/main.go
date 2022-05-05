package main

import "fmt"

//结构体构造函数：构造一个结构体实力的函数

type person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p1 := newPerson("辗迟", "侠岚", 18)
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}
